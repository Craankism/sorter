# 📄 sorter

`sorter` is a beginner-friendly command-line tool written in **Go** that sorts items from a text file. Each line in the input file is treated as an individual item. The program supports sorting in ascending or descending order and writing the result to a specified output file.

---

## ✅ Features

- Takes a text file as input (one item per line)
- Sorts lines in ascending or descending order
- Outputs sorted items to a new file
- Includes a `--debug` flag to show step-by-step sorting debug logs
- Implements a simple sorting algorithm (e.g., [bubble sort](https://studyflix.de/informatik/bubblesort-1325))

---

## 🛠 Usage

```bash
# Basic usage (ascending sort, default output path "output.txt")
go run main.go input.txt

# Sort in descending order
go run main.go input.txt --order desc

# Specify a custom output file
go run main.go input.txt --output sorted.txt

# Enable debug mode to show sorting steps
go run main.go input.txt --debug
```



## 🧾 Arguments
| Argument  | Type       | Required | Description                                     |
| --------- | ---------- | -------- | ----------------------------------------------- |
| input.txt | Positional | ✅ Yes    | Path to the input text file (one item per line) |
| --order   | Optional   | ❌ No     | asc (default) or desc for sorting order         |
| --output  | Optional   | ❌ No     | Path to output file. Default is output.txt      |
| --debug   | Flag       | ❌ No     | Enables debug logging of sorting steps          |



## 📥 Input Format
Input should be a plain text file with one item per line:

```
nginx
Copy
Edit
orange
apple
banana
```

## 📤 Output Format


Depending on sort order:

Ascending (asc):

```
nginx
Copy
Edit
apple
banana
orange
```

Descending (desc):

```
nginx
Copy
Edit
orange
banana
apple
```
Output is written to output.txt by default or the file provided via --output.



## Example

```bash
# With debug output
go run main.go items.txt --order asc --debug
```

Example debug output:

```bash
[DEBUG] Step 1: [orange apple banana]
[DEBUG] Step 2: [apple orange banana]
[DEBUG] Step 3: [apple banana orange]
```


🎓 Learning Goals
This project is intended for beginner software engineering students to learn:

- Basic Go CLI development
- File I/O operations
- Argument parsing with Go libraries (flag or cobra)
- Implementing and visualizing sorting algorithms
- Writing modular, testable code



## 🏗 Suggested Project Structure


```
sorter/
├── main.go
├── sorter/
│   ├── sorter.go      # Sorting logic
│   └── utils.go       # File I/O helpers
├── input.txt
├── output.txt
└── README.md
```


## 📎 Notes
Sorting uses a basic [bubble sort](https://studyflix.de/informatik/bubblesort-1325)to aid in learning.

Only plain text files are supported.

Debug logs are printed to standard output.


## Advanced Software Development Lifecycle Excercise

## 🚀 CI/CD and Release Automation
This repository uses GitHub Actions for continuous integration and automated release management.


## ✅ GitHub Actions Workflow
A GitHub Actions workflow is configured to:
- Build the Go CLI binary for your target platform(s)
- Upload the binary as a downloadable job artifact
- (Optional) Run unit tests if available


🏗 Example Workflow: `.github/workflows/build.yml`


```yml
name: Build CLI Binary

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'

      - name: Build sorter
        run: |
          mkdir -p dist
          go build -o dist/sorter ./main.go

      - name: Upload Binary Artifact
        uses: actions/upload-artifact@v4
        with:
          name: sorter-binary
          path: dist/sorter

```

> 💡 Once this is merged and running, navigate to a successful run under Actions tab to download the built binary as an artifact.

## 🏷️ Conventional Commits
We strictly encourage the use of [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) to enable automated versioning and changelog generation.

✅ Examples:

```
git commit -m "feat: add descending order support"
git commit -m "fix: correct file output bug"
git commit -m "chore: update README with CLI instructions"
```


## 🚀 Automated Releases with release-please
Once the build job is working, we use release-please to automate:

Version bumps using conventional commits

CHANGELOG.md generation

GitHub Releases


## 🧩 Example Workflow: .github/workflows/release.yml

```yml
name: Release

on:
  push:
    branches:
      - main

jobs:
  release:
    permissions:
      contents: write
      pull-requests: write

    runs-on: ubuntu-latest

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Release Please
        uses: google-github-actions/release-please-action@v4
        with:
          release-type: go
          package-name: sorter

```

Once configured:
- When a PR is merged with a proper commit message (e.g., feat: add feature), a release PR is automatically created.
- Once that PR is merged, a new GitHub Release is published and tagged.


## 🧪 Final Notes
- ✅ CI runs on every push and pull request to main
- 🎉 Releases are triggered automatically based on your commits
- 📦 Artifacts are available for download via GitHub Actions UI
- 💡 Use semantic versioning and conventional commits to drive clean releases