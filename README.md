# homepkg

## Project goals

* A package manager for you local scripts.
* List all available scripts.
* Execute your available scripts.
* Can sync scripts between computers. (Git clone, Get?)
* One file local configuration.
* Uses git repositories as source. (Public & private)
* Every repo can contain a single project or multiple projects.
* You can choose which scripts you want from a repo and do not need to checkout all scripts.
* Needs to be able to run install scripts.
* Add & remove scripts.
* Commands can be grouped.
* Supports a local folder.
* Creates symlinks for bin folder which you can add to your $PATH so that all scripts are available.

## Project values
* No silent errors!
* Fully transparent for the user.

## Technical details
* Uses a YAML configuration.
* Uses your command line git.
* You have one main configuration file on your computer which specifies which script you use from which repositories.
* Every git repo can have a configuration in its root to specify which packages exist in the repo.
* Every script can have a configuration file to specify an install script.
* Project configuration can container a group.

* The repos are cloned into `~/.homepkg`.

* Directory structure:
```sh
~/.homepkg/
    bin/ -> contains symlinks to all packages. Can be added to $PATH
    repo-1
    .
    .
    .
    repo-N

~/my-scripts/hello.sh -> ~/.homepkg/bin/hello.sh (contains `homepkg.yaml`)
~/my-scripts/script1/world.sh -> ~/.homepkg/bin/world.sh
~/other-scripts/hello-world.sh -> ~/.homepkg/bin/hello-world.sh
```

* Layout `homepkg.yaml` or `homepkg.yml`:
```yaml
homepkg:
  gitUrl: https://github.com/alice/my-scripts
  author: Alice
  packages:
    - name: Package-1
      author: Alice
      authorMail: alice@alice.com
      description: Best script!
      path: script/
      command: hello.sh # Must be a single file.
      install: yarn install # Can be any command that will be executed within the package directory.
    - name: Package-2
      author: Bob
      authorMail: bob@bob.com
      description: My best script!
      path: script-2/
      command: world.sh
      install: npm install
```

### Commands

#### add
* Installs **all** commands defined in a `homepkg.yaml`.
* If a repository has already been added the `update` command will be executed.

#### update
1. (For remote repos: `git pull`)
For each package:
    1. Delete symlinks
    2. Run `install` step
        * If this step fails homepkg should create a dummy command that is named like original command of the package.
          The dummy command will print an error message with the log output of the failed installation step.
    3. Create symlinks (Needed to account for changed package and command names)

#### remove
* Deletes remote repositories
* Deletes symlinks

#### list
* Show all local and remote packages in `~/.homepkg/bin`

```
homepkg:

Commands:
- add - Clones the script from its repo and runs the install script
- list - Lists all scripts
- remove - Let's you remove a script, let's you remove the script from
- update - Updates a repository.

otherwise run
```
