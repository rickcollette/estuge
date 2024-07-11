package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/russross/blackfriday/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

const outputDir = "static"

type Config struct {
    SiteName            string `yaml:"SiteName"`
    SiteDescription     string `yaml:"SiteDescription"`
    SiteTitle           string `yaml:"SiteTitle"`
    FontSet             string `yaml:"FontSet"`
    NumberOfSlides      int    `yaml:"NumberOfSlides"`
    SlideChangeInterval int    `yaml:"SlideChangeInterval"`
    SlideFolder         string `yaml:"SlideFolder"`
    SlidePrefix         string `yaml:"SlidePrefix"`
    SocialX             string `yaml:"SocialX"`
    SocialFB            string `yaml:"SocialFB"`
    SocialYT            string `yaml:"SocialYT"`
}

type Page struct {
    Title   string
    Header  template.HTML
    Content template.HTML
    Footer  template.HTML
    Config  Config
}

// Helper functions for template
var funcMap = template.FuncMap{
    "split":         strings.Split,
    "trimExtension": trimExtension,
    "title":         toTitle,
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: estuge <input-folder> [template-name]\nVisit https://govamp.org for information on how to use Estuge.")
        os.Exit(1)
    }

    inputDir := os.Args[1]
    templateName := "base"
    if len(os.Args) >= 3 {
        templateName = os.Args[2]
    }

    templateDir := filepath.Join("templates", templateName)

    // Read the configuration file
    configPath := filepath.Join(templateDir, "config.yaml")
    configFile, err := os.ReadFile(configPath)
    if err != nil {
        fmt.Printf("Error reading config file: %v\n", err)
        os.Exit(1)
    }

    var config Config
    err = yaml.Unmarshal(configFile, &config)
    if err != nil {
        fmt.Printf("Error parsing config file: %v\n", err)
        os.Exit(1)
    }

    // Read the template files
    mainTemplatePath := filepath.Join(templateDir, "main.tmpl")
    headerTemplatePath := filepath.Join(templateDir, "header.tmpl")
    footerTemplatePath := filepath.Join(templateDir, "footer.tmpl")
    cssPath := filepath.Join(templateDir, "style", "style.css")
    jsDir := filepath.Join(templateDir, "js")
    imagesDir := filepath.Join(templateDir, "images")

    mainTmplContent, err := os.ReadFile(mainTemplatePath)
    if err != nil {
        fmt.Printf("Error reading main template file: %v\n", err)
        os.Exit(1)
    }

    headerTmplContent, err := os.ReadFile(headerTemplatePath)
    if err != nil {
        fmt.Printf("Error reading header template file: %v\n", err)
        os.Exit(1)
    }

    footerTmplContent, err := os.ReadFile(footerTemplatePath)
    if err != nil {
        fmt.Printf("Error reading footer template file: %v\n", err)
        os.Exit(1)
    }

    tmpl, err := template.New("page").Funcs(funcMap).Parse(string(mainTmplContent))
    if err != nil {
        fmt.Printf("Error parsing main template file: %v\n", err)
        os.Exit(1)
    }

    // Copy CSS, JS, and images to static directory
    copyStaticFiles(cssPath, "static/css/style.css")
    copyStaticFiles(jsDir, "static/js")
    copyStaticFiles(imagesDir, "static/images")

    // Build the navigation
    nav, err := buildNavigation(inputDir)
    if err != nil {
        fmt.Printf("Error building navigation: %v\n", err)
        os.Exit(1)
    }

    // Parse the header template with navigation
    header, err := template.New("header").Funcs(funcMap).Parse(string(headerTmplContent))
    if err != nil {
        fmt.Printf("Error parsing header template file: %v\n", err)
        os.Exit(1)
    }
    var headerContent strings.Builder
    headerPage := Page{Config: config}
    if err := header.Execute(&headerContent, nav); err != nil {
        fmt.Printf("Error executing header template: %v\n", err)
        os.Exit(1)
    }

    headerHTML := template.HTML(headerContent.String())

    // Parse the footer template with the config
    footer, err := template.New("footer").Funcs(funcMap).Parse(string(footerTmplContent))
    if err != nil {
        fmt.Printf("Error parsing footer template file: %v\n", err)
        os.Exit(1)
    }
    var footerContent strings.Builder
    if err := footer.Execute(&footerContent, headerPage); err != nil {
        fmt.Printf("Error executing footer template: %v\n", err)
        os.Exit(1)
    }

    footerHTML := template.HTML(footerContent.String())

    err = filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
        return processFile(path, info, err, tmpl, inputDir, headerHTML, footerHTML, config)
    })
    if err != nil {
        fmt.Printf("Error processing files: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Conversion complete!")
}

// toTitle converts a string to title case using the new golang.org/x/text/cases package
func toTitle(s string) string {
    c := cases.Title(language.Und)
    return c.String(s)
}

// trimExtension removes the .html extension from a string
func trimExtension(s string) string {
    return strings.TrimSuffix(s, ".html")
}

func buildNavigation(rootDir string) (map[string][]string, error) {
    nav := make(map[string][]string)

    err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.IsDir() {
            return nil
        }

        if strings.HasSuffix(info.Name(), ".md") {
            relPath, err := filepath.Rel(rootDir, path)
            if err != nil {
                return err
            }

            parts := strings.Split(relPath, string(os.PathSeparator))
            if len(parts) == 1 {
                nav["Home"] = append(nav["Home"], "/")
            } else {
                dir := parts[0]
                file := strings.TrimSuffix(parts[len(parts)-1], ".md")
                if file == "index" {
                    nav[dir] = append(nav[dir], "/"+dir+"/")
                } else {
                    nav[dir] = append(nav[dir], "/"+dir+"/"+file+"/")
                }
            }
        }

        return nil
    })

    return nav, err
}

func processFile(path string, info os.FileInfo, err error, tmpl *template.Template, inputDir string, header template.HTML, footer template.HTML, config Config) error {
    if err != nil {
        return err
    }

    if info.IsDir() {
        return nil
    }

    if strings.HasSuffix(info.Name(), ".md") {
        return convertMarkdownToHTML(path, tmpl, inputDir, header, footer, config)
    }

    return nil
}

func convertMarkdownToHTML(inputPath string, tmpl *template.Template, inputDir string, header template.HTML, footer template.HTML, config Config) error {
    content, err := os.ReadFile(inputPath)
    if err != nil {
        return err
    }

    relativePath, err := filepath.Rel(inputDir, inputPath)
    if err != nil {
        return err
    }

    // Convert Markdown to HTML
    htmlContent := blackfriday.Run(content)
    
    // Process custom CSS class and attribute syntax
    customClassRegex := regexp.MustCompile(`(.*?)\{:\s*\.([a-zA-Z0-9_-]+)\s*\}`)
    customAttrRegex := regexp.MustCompile(`(.*?)\{:\s*\#([a-zA-Z0-9_-]+)\s*\}`)
    processedContent := customClassRegex.ReplaceAll(htmlContent, []byte(`<div class="$2">$1</div>`))
    processedContent = customAttrRegex.ReplaceAll(processedContent, []byte(`<div id="$2">$1</div>`))

    title := filepath.Base(inputPath)
    title = strings.TrimSuffix(title, filepath.Ext(title))

    page := Page{
        Title:   title,
        Header:  header,
        Content: template.HTML(processedContent),
        Footer:  footer,
        Config:  config,
    }

    outputPath := determineOutputPath(relativePath)
    err = os.MkdirAll(filepath.Dir(outputPath), 0755)
    if err != nil {
        return err
    }

    outputFile, err := os.Create(outputPath)
    if err != nil {
        return err
    }
    defer outputFile.Close()

    return tmpl.Execute(outputFile, page)
}

func determineOutputPath(relativePath string) string {
    baseName := strings.TrimSuffix(filepath.Base(relativePath), ".md")
    if baseName == "index" {
        return filepath.Join(outputDir, strings.TrimSuffix(relativePath, ".md")+".html")
    } else {
        outputDirPath := filepath.Join(outputDir, filepath.Dir(relativePath), baseName)
        return filepath.Join(outputDirPath, "index.html")
    }
}

func copyStaticFiles(src, dest string) error {
    info, err := os.Stat(src)
    if err != nil {
        return err
    }

    if info.IsDir() {
        return copyDir(src, dest)
    } else {
        return copyFile(src, dest)
    }
}

func copyDir(src, dest string) error {
    entries, err := os.ReadDir(src)
    if err != nil {
        return err
    }

    err = os.MkdirAll(dest, 0755)
    if err != nil {
        return err
    }

    for _, entry := range entries {
        srcPath := filepath.Join(src, entry.Name())
        destPath := filepath.Join(dest, entry.Name())

        if entry.IsDir() {
            if err := copyDir(srcPath, destPath); err != nil {
                return err
            }
        } else {
            if err := copyFile(srcPath, destPath); err != nil {
                return err
            }
        }
    }

    return nil
}

func copyFile(src, dest string) error {
    input, err := os.ReadFile(src)
    if err != nil {
        return err
    }

    err = os.MkdirAll(filepath.Dir(dest), 0755)
    if err != nil {
        return err
    }

    return os.WriteFile(dest, input, 0644)
}
