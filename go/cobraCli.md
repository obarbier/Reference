## Why and why not?
-------------------------
Let's start with this. Why should one use go to create a CLI application

- Cross-Platform Support
  - Go’s cross-platform compilation gives us the benefits of a compiled language (reduced dependencies, etc), with the portability of a scripting language.
- Distribution Flexibility
  - but since Go is a compiled language, we have a number of options for distributing our applications. The ability to download a single file and execute commands - without an installer, or even setup process - does wonders to a user.

why not ?
- No one a team can claim to have "deep experience" with go
- Hard to process JSON data becasue of extreme type safety

## Getting started
--------------------
Using [Hoarder](https://github.com/nanopack/hoarder) and [Cobra](https://github.com/spf13/cobra) as rdocumentation and an example. While you are welcome to provide your own organization, typically a Cobra-based application will follow the following organizational structure:

```
▾ appName/
  ▾ cmd/
      add.go
      your.go
      commands.go
      here.go
    main.go
```
The commands.go file is the meat of the CLI, all other files in the package represent actual commands for the CLI. Cobra provides its [own program](https://github.com/spf13/cobra/blob/master/cobra/README.md) that will create your application and add any commands you want. It's the easiest way to incorporate Cobra into your application.

#### Viper

Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats. It supports:
- setting defaults
- reading from JSON, TOML, YAML, HCL, and Java properties config files
- live watching and re-reading of config files (optional)
- reading from environment variables
- reading from remote config systems (etcd or Consul), and watching changes
- reading from command line flags
- reading from buffer
- setting explicit values

Viper does the following for you:

- Find, load, and unmarshal a configuration file in JSON, TOML, YAML, HCL, or Java properties formats.
- Provide a mechanism to set default values for your different configuration options.
- Provide a mechanism to set override values for options specified through command line flags.
- Provide an alias system to easily rename parameters without breaking existing code.
- Make it easy to tell the difference between when a user has provided a command line or config file which is the same as the default.

## Cli Description
----------------------------
We will create a Todo CLI app in go it will perform the following
1. Create to do : `--create -c <string>`
2. get all todo : `--getAll -G`
3. get todo :`--get -g <id>`
4. Delete to do : `--delete -d <id>`
5. Delete all to do : `--deleteAll -D `
6. Update todo : `--update -u <id> <string>`
7. get status : `--getStatus -s`
8. tag todo :

## Let's code
----------------
1. Cobra init
  1. cobra init todoCLI --pkg-name todoCLI
  2. Make changes in root.go and main.go.
  3. test
2. Set GOPATH, GOROOT, GOBIN
3. install github.com/mattn/go-sqlite3
  1. Had to fix issue by installing mingw-w64
  2. setup sqlite in init code
4. Adding create task
  1. set up Viper

## reference
------------------
1. [dev.to](https://dev.to/uilicious/why-we-migrated-our-cli-from-nodejs-to-golang-1ol8)
2. [yext](http://engblog.yext.com/post/going-all-in-with-go-for-cli-apps)
3. [medium](https://medium.com/@skdomino/writing-better-clis-one-snake-at-a-time-d22e50e60056)
4. [GoPath vs GoRoot](https://stackoverflow.com/questions/7970390/what-should-be-the-values-of-gopath-and-goroot)
5. [mingw-w64](https://sourceforge.net/projects/mingw-w64/files/latest/download)
6. [Windows: "gcc": executable file not found in %PATH% ](https://github.com/mattn/go-sqlite3/issues/467)
