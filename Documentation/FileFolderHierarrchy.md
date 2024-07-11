# Understanding the Folder Structure for estuge 
## (S2G - Static Site Generator)

This document explains the folder structure required for estuge to work correctly. It details what files and folders need to be created and where to place your content.

## Project Structure

Your project directory should be organized as follows:

**NOTE:** The `my_project` and `mysite/` directories, and the files and folders below are used as an example. You can name it whatever you like.

```
my_project/
├── templates/
│       ├── config.yaml
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
```

### Root Directory (`mysite/`)

The root directory of your project contains the main content of your site. The root directory should include:

- `index.md`: The main homepage of your site.

### `about/` Directory

The `about/` directory contains information about your site or organization. This directory should include:

- `index.md`: The main content for the "About" section of your site.

### `contact/` Directory

The `contact/` directory contains contact information or a contact form for your site. This directory should include:

- `index.md`: The main content for the "Contact" section of your site.

### `docs/` Directory

The `docs/` directory is for documentation or other similar content. This directory should include:

- `index.md`: The main content for the "Docs" section of your site.
- `other.md`: Additional content or documentation related to the "Docs" section.  When estuge  reads this file, it will generate a new folder called other, and the output will be an index.html file inside the other folder.

## Example Markdown Files

### `index.md` (Root Directory)

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

### `about/index.md`

```markdown
# About Us

This page provides information about our site or organization.

## Our Mission

Our mission is to provide high-quality content and services.

## Team

Meet our dedicated team.
```

### `contact/index.md`

```markdown
# Contact Us

Feel free to reach out to us through the following contact methods.

## Email

You can email us at [contact@example.com](mailto:contact@example.com).

## Social Media

Follow us on social media:
- Twitter: [@example](https://twitter.com/example)
- Facebook: [Example](https://facebook.com/example)
```

### `docs/index.md`

```markdown
# Documentation

Welcome to the documentation section. Here you will find all the information you need to use our site or service.

## Getting Started

To get started, follow these steps.

## FAQ

Find answers to frequently asked questions.
```

### `docs/other.md`

```markdown
# Other Documentation

This section contains additional documentation.

## Advanced Topics

Learn about advanced topics and features.
```

## Running the Static Site Generator

1. **Navigate to your project directory:**

    ```bash
    cd mysite
    ```

2. **Build and run the application:**

    ```bash
    estuge mysite base
    ```

    This command will process the Markdown files in the `mysite/` directory using the `base` template and generate the static site in the `static/` directory.

3. **Copy the contents of the static/ folder to your web server.**

    ```bash
    cp -r static/* /var/www/html/
    ```

    This command will copy the contents of the `static/` directory to your web server's document root.

4. **Open the generated site in your browser:**

It's as easy as that.


## Customizing Your Site

You can customize your site by modifying the template files in the `templates/base/` directory. The structure of these files determines how the content is rendered on your site.

By following this guide, you can organize your content and structure your project directory correctly to use the static site generator effectively. Happy coding!
