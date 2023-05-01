# trans.nvim

gpt_trans.nvim is a plugin for Neovim to translate English to Japanese with GPT3.5-tarbo.

## Required

You must first set up authentication by creating a API Key or service account of GCP.

Then you need to install the [Go](https://go.dev/doc/install).

Author using the following version.
```
$ go version
go version go1.20 linux/amd64
```

### Using API Key

The API Key documentation can be found [here](https://platform.openai.com/docs/guides/production-best-practices/api-keys).

Please set the environment variable `OPENAI_API_KEY` to the API Key.

## Note

You need a little bit of money to use OpenAI API.

For more information [here](https://openai.com/pricing).

## Installation

For vim-plug
```viml
Plug 'devoc09/gpt_trans.nvim', {'do': 'make'}
```

## Settings

```viml
let g:trans_lang_output = 'preview'
```

### To use floating windows

Floating windows are supported in neovim >= 0.4.0.

```viml
let g:trans_lang_output = 'float'
```

A Floating window is automatically hide when cursor is moved.
