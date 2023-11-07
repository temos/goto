# GOTO

A simple command-line utility to quickly find and switch to project folders

## Example Usage
```shell
# add to.bashrc:
function goto {
    $to=(\goto \
        "/home/user/code/go" GO \
        "/home/user/code/dotnet" DOTNET \
        "/home/user/projects" PROJECTS) && cd "$to"
}

# run in shell:
goto
```