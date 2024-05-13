# freqli

freqli is a command-line tool that analyzes command frequency in your shell history files. It helps you identify the most commonly used commands and optimize your workflow.

## Features

- Analyze command frequency in Bash and Zsh history files.
- Supports both Bash and Zsh shells.
- Easy to install and use.

## Installation

To install freqli, you can use the `go install` command:

```sh
go install github.com/Shobhit-Nagpal/freqli@latest
```

Make sure you have Go installed on your system. You can download and install it from [https://golang.org/dl/](https://golang.org/dl/).

## Usage

After installing freqli, you can use it to analyze your shell history files. Here's how you can use it:

```sh
freqli --shell <shell_type>
```

Replace `<shell_type>` with either `bash` or `zsh`, depending on the type of shell you're using.

For example, to analyze your Bash history:

```sh
freqli --shell bash
```

And to analyze your Zsh history:

```sh
freqli --shell zsh
```

## Contributing

Contributions are welcome! If you have any ideas for new features, improvements, or bug fixes, feel free to open an issue or submit a pull request.

## Acknowledgments

Credits to [Boot.dev](https://boot.dev) for introducing the history command in their course.
