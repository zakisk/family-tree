## family-tree
This project is a command-line interface (CLI) application developed in Go, leveraging the powerful Cobra library. The primary objective of this application is to facilitate the storage and management of family tree information. Utilizing the sqlite3 database library ensures a reliable foundation for storing and retrieving crucial details. The application features a user-friendly interface, providing a seamless experience for users to organize and access essential information within their family tree.

## Example
<img src="/tree.gif" height="500px"/>

## family-tree add

adds names and roles

### Synopsis

Subcommand to add names of family members and roles of them. 
For example:
// adds a person in family
family-tree add person john

// adds an role entry e.g father, mother, son, daughter
family-tree add role father


```
family-tree add [flags]
```

### Options

```
  -h, --help   help for add
```

### SEE ALSO

* [family-tree](family-tree.md)	 - family tree utility CLI app
* [family-tree add person](family-tree_add_person.md)	 - adds person
* [family-tree add role](family-tree_add_role.md)	 - adds role

###### Auto generated by spf13/cobra on 27-Aug-2023
## family-tree add person

adds person

### Synopsis

Subcommand to add persons (duplicate not allowed). 
For example:
// adds a person in family
family-tree add person [person-name]


```
family-tree add person [flags]
```

### Options

```
  -h, --help   help for person
```

### SEE ALSO

* [family-tree add](family-tree_add.md)	 - adds names and roles

###### Auto generated by spf13/cobra on 27-Aug-2023
## family-tree add role

adds role

### Synopsis

Subcommand to add roles (duplicate not allowed). 
For example:
// adds role of a person
// allowed roles [father, mother, son, daughter, brother, sister, wife, husband]
family-tree add role [role-name]


```
family-tree add role [flags]
```

### Options

```
  -h, --help   help for role
```

### SEE ALSO

* [family-tree add](family-tree_add.md)	 - adds names and roles

###### Auto generated by spf13/cobra on 27-Aug-2023
## family-tree completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(family-tree completion bash)

To load completions for every new session, execute once:

#### Linux:

	family-tree completion bash > /etc/bash_completion.d/family-tree

#### macOS:

	family-tree completion bash > $(brew --prefix)/etc/bash_completion.d/family-tree

You will need to start a new shell for this setup to take effect.


```
family-tree completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [family-tree completion](family-tree_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 27-Aug-2023
## family-tree completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	family-tree completion fish | source

To load completions for every new session, execute once:

	family-tree completion fish > ~/.config/fish/completions/family-tree.fish

You will need to start a new shell for this setup to take effect.


```
family-tree completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [family-tree completion](family-tree_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 27-Aug-2023
## family-tree completion

Generate the autocompletion script for the specified shell

### Synopsis

Generate the autocompletion script for family-tree for the specified shell.
See each sub-command's help for details on how to use the generated script.


### Options

```
  -h, --help   help for completion
```

### SEE ALSO

* [family-tree](family-tree.md)	 - family tree utility CLI app
* [family-tree completion bash](family-tree_completion_bash.md)	 - Generate the autocompletion script for bash
* [family-tree completion fish](family-tree_completion_fish.md)	 - Generate the autocompletion script for fish
* [family-tree completion powershell](family-tree_completion_powershell.md)	 - Generate the autocompletion script for powershell
* [family-tree completion zsh](family-tree_completion_zsh.md)	 - Generate the autocompletion script for zsh

###### Auto generated by spf13/cobra on 27-Aug-2023
## family-tree completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	family-tree completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
family-tree completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [family-tree completion](family-tree_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 27-Aug-2023
## family-tree completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(family-tree completion zsh)

To load completions for every new session, execute once:

#### Linux:

	family-tree completion zsh > "${fpath[1]}/_family-tree"

#### macOS:

	family-tree completion zsh > $(brew --prefix)/share/zsh/site-functions/_family-tree

You will need to start a new shell for this setup to take effect.


```
family-tree completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [family-tree completion](family-tree_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 27-Aug-2023
## family-tree connect

connects two person

### Synopsis

Subcommand to connect two persons with a relationship.
For example:
// connects two persons of a family
family-tree connect [person1] --as brother --of [person2]

// default gender of person2 is male if it isn't specify it with --person2-gender or -g
family-tree connect [person1] --as brother --of [person2] --person2-gender female


```
family-tree connect [flags]
```

### Options

```
  -a, --as string               specifies the role of the person
  -h, --help                    help for connect
  -o, --of string               specifies the relation with person2
  -g, --person2-gender string   specifies gender of person2 (default "male")
```

### SEE ALSO

* [family-tree](family-tree.md)	 - family tree utility CLI app

###### Auto generated by spf13/cobra on 27-Aug-2023
## family-tree count

counts relatives

### Synopsis

Subcommand to count relatives by their roles.
For example:
// count sons
family-tree count sons --of [person]

// count daughters
family-tree count daughters --of [person]


```
family-tree count [flags]
```

### Options

```
  -h, --help                   help for count
  -o, --of string              specifies the person
  -g, --person-gender string   specifies person gender (default "male")
```

### SEE ALSO

* [family-tree](family-tree.md)	 - family tree utility CLI app

###### Auto generated by spf13/cobra on 27-Aug-2023
## family-tree father

shows father name

### Synopsis

Subcommand to find father name.
For example:
family-tree father --of [person]


```
family-tree father [flags]
```

### Options

```
  -h, --help        help for father
  -o, --of string   specifies person name
```

### SEE ALSO

* [family-tree](family-tree.md)	 - family tree utility CLI app

###### Auto generated by spf13/cobra on 27-Aug-2023
## family-tree

family tree utility CLI app

### Synopsis

An utility tool to create and view family tree records.

```
family-tree [flags]
```

### Options

```
  -h, --help   help for family-tree
```

### SEE ALSO

* [family-tree add](family-tree_add.md)	 - adds names and roles
* [family-tree completion](family-tree_completion.md)	 - Generate the autocompletion script for the specified shell
* [family-tree connect](family-tree_connect.md)	 - connects two person
* [family-tree count](family-tree_count.md)	 - counts relatives
* [family-tree father](family-tree_father.md)	 - shows father name

###### Auto generated by spf13/cobra on 27-Aug-2023
