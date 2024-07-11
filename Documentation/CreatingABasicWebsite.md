# Creating a Website Using the Static Site Generator

This guide will help you create a website using the static site generator application. The application converts Markdown files into HTML files using customizable templates.

## Prerequisites

- Go programming language installed on your system.
- Basic knowledge of Markdown.
- A text editor.

## Project Structure

Your project should follow this structure:

```
my_project/
├── config.yaml
├── templates/
│   └── base/
│       ├── footer.tmpl
│       ├── header.tmpl
│       ├── main.tmpl
│       ├── images/
│       ├── js/
│       │   └── helper_funcs.js
│       └── style/
│           └── style.css
├── mysite/
│   ├── index.md
│   ├── about/
│   │   └── index.md
│   ├── contact/
│   │   └── index.md
│   └── docs/
│       ├── index.md
│       └── other.md
└── static/
    ├── css/
    ├── images/
    ├── js/
    └── ...
```

## Configuration File

Create a `config.yaml` file inside the `templates/base/` directory with the following content:

```yaml
SiteName: "MySite"
SiteDescription: "This is my site"
SiteTitle: "MySite"
FontSet: "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css" # FontAwesome CDN

# Slider Info
NumberOfSlides: 3
SlideChangeInterval: 3000 # Milliseconds
SlideFolder: "images/slides"
SlidePrefix: "slide_" # Results in slide_[1...${NumberOfSlides}].jpg

# Footer Info
SocialX: "https://twitter.com/youraccount"
SocialFB: "https://facebook.com/youraccount"
SocialYT: "https://youtube.com/youraccount"
```

## Template Files

### `header.tmpl`

```html
[HEADER.TMPL HERE]
```

### `footer.tmpl`

```html
[FOOTER.TMPL HERE]
```

### `main.tmpl`

```html
[MAIN.TMPL HERE]
```

### `style.css`

```css
[STYLE.CSS HERE]
```

## JavaScript File

### `helper_funcs.js`

```javascript
[HELPER_FUNCS.JS HERE]
```

## Markdown Files

Create your content in Markdown files inside the `mysite/` directory.

### Example Markdown File: `mysite/index.md`

```markdown
# Welcome to My Site

![My Image](images/image.png){:.img-class}

This is the home page of my static site generated from Markdown files.

## Custom Section

This section has a custom CSS class.

{:.custom-class}
### Subsection

This subsection also has a custom CSS class.
```

## Running the Application

1. **Navigate to your project directory:**

    ```bash
    cd my_project
    ```

2. **Build and run the application:**

    ```bash
    go run main.go mysite base
    ```

    This command will process the Markdown files in the `mysite/` directory using the `base` template and generate the static site in the `static/` directory.

3. **Open the generated site in your browser:**

    ```bash
    open static/index.html
    ```

## Customizing Your Site

You can customize your site by modifying the template files in the `templates/base/` directory. The structure of these files determines how the content is rendered on your site.

- **header.tmpl**: Defines the header section of your site, including navigation links.
- **footer.tmpl**: Defines the footer section of your site, including social media links with Font Awesome icons.
- **main.tmpl**: Defines the main structure of your HTML pages, including the header, footer, and content sections.
- **style.css**: Defines the styles for your site, including basic resets, header, footer, dropdown menus, and slider styles.
- **helper_funcs.js**: Includes JavaScript functionality for dropdown menus and image sliders.

## Adding Custom CSS Classes and Attributes in Markdown

You can add custom CSS classes and attributes to your Markdown content using the following syntax:

### Example Markdown with Custom Classes and Attributes

```markdown
# Welcome to My Site

![My Image](images/image.png){:.img-class}

This is the home page of my static site generated from Markdown files.

## Custom Section

This section has a custom CSS class.

{:.custom-class}
### Subsection

This subsection also has a custom CSS class.
```

In the above example:
- `{: .img-class}` adds a custom class to the image.
- `{: .custom-class}` adds a custom class to the section.

These custom classes will be processed and rendered correctly in the generated HTML files.

## Conclusion

By following this guide, you can create a static website using Markdown and customizable templates. Modify the configuration and template files as needed to fit your specific requirements. Happy coding!
