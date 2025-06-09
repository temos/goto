# GOTO

A simple command-line utility to quickly find and switch to project folders

## Install

```shell
go install -trimpath -ldflags="-s -w" github.com/temos/goto@latest
```

## Usage

goto is configured using a yaml configuration file, the path to which has to be passed as the first argument:
```shell
goto ./config.yml
```

configuration file example:
```yaml
activeColor: "#8C18E2" # optional, the color of the currently selected menu entry
directories:
  - prefix: GO
    path: "C://source/go"
  - prefix: DOTNET
    path: "C://source/dotnet"
  - prefix: PROJECTS
    path: "C://projects"
urls:
  - prefix: DOCKER
    name: HUB
    url: https://hub.docker.com/
```

goto returns a non-zero exit code when it is terminated by pressing **ctrl+c**

## Example Usage

### Bash
//TODO: validate that this works
```bash
# add to .bashrc:
goto_func() {
    to=$(\goto config.yml)
    if [ ! $? -eq 0 ]; then
        return
    fi

    if [[ "$to" == *"://"* ]]
        xdg-open "$to"
    else
        cd "$to"
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
  $to=(goto.exe . config.yml)
  if (!$?) {
    return
  }

  if ($to.Contains("://")) {
    start "$to"
  } else {
    cd "$to"
  }
}

# run in shell:
goto
