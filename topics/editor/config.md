## Definition
Neovim allows for multiple configuration managenemt. Packer and Plug. based on [here](https://www.reddit.com/r/neovim/comments/qexmy6/plug_vs_packer/?utm_source=BD&utm_medium=Search&utm_name=Bing&utm_content=PSR1) looks like packer has some performance benefits 

### Packer quick start
```
# install
git clone --depth 1 https://github.com/wbthomason/packer.nvim\
 ~/.local/share/nvim/site/pack/packer/start/packer.nvim
```

```
# setting plugins under ~/.config/nvim/lua/plugins.lua

```


the commad for packer are as follow
```
-- You must run this or `PackerSync` whenever you make changes to your plugin configuration
-- Regenerate compiled loader file
:PackerCompile

-- Remove any disabled or unused plugins
:PackerClean

-- Clean, then install missing plugins
:PackerInstall

-- Clean, then update and install plugins
-- supports the `--preview` flag as an optional first argument to preview updates
:PackerUpdate

-- Perform `PackerUpdate` and then `PackerCompile`
-- supports the `--preview` flag as an optional first argument to preview updates
:PackerSync

-- Loads opt plugin immediately
:PackerLoad completion-nvim ale
```


## Glossary
* **suspendisse**: dictum at tempor commodo ullamcorper a lacus vestibulum sed [LINK](https://loremipsum.io/generator/?n=9&t=w)

## Reference
* [Packer Github Page](https://github.com/wbthomason/packer.nvim#quickstart)
* [Lua + Packer example](https://gist.github.com/benfrain/97f2b91087121b2d4ba0dcc4202d252f)
* [Lua neovim guide](https://github.com/nanotee/nvim-lua-guide)
* [more](https://icyphox.sh/blog/nvim-lua/)