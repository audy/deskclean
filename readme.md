# Deskclean [![Go](https://github.com/audy/deskclean/actions/workflows/go.yml/badge.svg)](https://github.com/audy/deskclean/actions/workflows/go.yml)

Keep your `~/Desktop` tidy!

Deskclean moves files on your `~/Desktop` to sub-directories based on their file extensions.

# Installation

Grab a binary from the [latest release](https://github.com/audy/deskclean/releases/latest)

## Usage

```sh
$ deskclean
```

Run `deskclean -h` for help:

```sh
Usage of ./deskclean:
  -path string
        path to clean (default "$HOME/Desktop")
```

## Config

Deskclean can be configured via `~/.config/deskclean/config.toml`. By default,
Deskclean uses the following config:

```toml
[textfiles]
pattern = ".*\\.(rtf|rtfd|md|txt|docx?|rtf|html?|pdf|log)$"

[data]
pattern = ".*\\.(ab1|csv|sam|fasta|fastq|fa|fna|faa|gbk?|gbf|gff|numbers|aln|zip|tar.gz|xlsx?|sqlite|json?)$"

[scripts]
pattern = ".*\\.(rmd|go|sql|pl|py|sh|rb|js|ts|coffee|c|r|R|ipynb)$"

[images]
pattern = ".*\\.(svg|jpe?g|png|gif|gifv|bmp|mp4|mov|m4v|ai|webp)$"
```

In this example, files that end in `.pdf` will be moved to
`~/Desktop/textfiles`, etc...
