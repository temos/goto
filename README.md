# GOTO

A simple command-line utility to quickly find and switch to project folders

## Usage

goto is configured by passing pairs of paths and display names as arguments:
```shell
goto [path1] [name1] [path2] [name2] ...
```

for example:
```shell
goto /home/user/code/go GO /home/user/code/dotnet DOTNET
```

goto returns a non-zero exit code when it is terminated by pressing **ctrl+c** or **q**

## Example Usage

### Bash
```bash
# add to .bashrc:
goto_func() {
    to=$(\goto \
        "/home/user/code/go" GO \
        "/home/user/code/cs" CSHARP)
        if [ $? -eq 0 ]; then
        cd "$to"
    else
        echo "$to"
    fi
}
alias goto=goto_func

# run in shell:
goto
```

### Powershell
```powershell
# add to powershell profile:
function goto {
  $to=(goto.exe C:\Users\user\Desktop\code\go\ GO C:\Users\user\Desktop\code\cs\ CSHARP)
  if ($?) { cd $to }
}

# run in shell:
goto
