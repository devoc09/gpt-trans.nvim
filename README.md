# trans.nvim

gpt_trans.nvim is a plugin for Neovim to translate English to Japanese with GPT3.5-tarbo.

![gpt-trans nvim](https://user-images.githubusercontent.com/50615605/235466220-99a9aab9-fdfd-48e1-a749-c5ed2d41c8ca.gif)

## Required
Install the [Go](https://go.dev/doc/install).

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
Plug 'devoc09/gpt-trans.nvim', {'do': 'make'}
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

## Thanks
- Inspired by [utahta/trans.nvim](https://github.com/utahta/trans.nvim).
