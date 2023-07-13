# HTML-Translator
HTML Translator is a Go program that allows you to translate HTML files into different languages using the Google Cloud Translate API.
## Prerequisites
Before using HTML Translator, make sure you have the following prerequisites:

- Go installed on your machine
- Access to the Google Cloud Translate API
- An HTML file to translate
- A service account JSON key file for authenticating with the Google Cloud Translate API

## Installation
1. Clone the repo
<br /> - ```git clone github.com/ahyeonn/HTML-Translator```
3. Cd - into project directory
<br/> - ```cd html-translator```
4. Build Go program
<br/> - ```go build```
## Usage
To translate an HTML file using HTML Translator, follow these steps:
1. Make sure you have the HTML file you want to translate available.
2. Run this command
<br/> - ```./translate -convert=path/to/your/file.html```
3. You will be prompted to choose the target language for translation. Enter the language code (e.g., "es" for Spanish, "fr" for French, etc.) and press Enter.

&nbsp;![image](https://github.com/Ahyeonn/HTML-Translator/assets/57298293/adee9e2e-2f03-486e-b946-36e86feeed7c)&nbsp;

5. The translated HTML file will be generated and saved as "translate.html" in the current directory.
## From English (en) to Spanish (es)
- <b>English Version</b>
<br/> ![image](https://github.com/Ahyeonn/HTML-Translator/assets/57298293/cccec675-581e-4c50-87ce-16ee2e15f1ab) <br/>
- <b>Spanish Version</b><br/>
![image](https://github.com/Ahyeonn/HTML-Translator/assets/57298293/507db554-220f-47e6-9624-6e897cb28cba)