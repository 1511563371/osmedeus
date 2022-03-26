# Osmedeus Core Engine

<p align="center">
  <img alt="Osmedeus" src="https://raw.githubusercontent.com/osmedeus/assets/main/logo-transparent.png" height="140" />
  <br />
  <strong>Osmedeus - A Workflow Engine for Offensive Security</strong>

  <p align="center">
  <a href="https://docs.osmedeus.org/"><img src="https://img.shields.io/badge/Documentation-0078D4?style=for-the-badge&logo=GitBook&logoColor=39ff14&labelColor=black&color=black"></a>
  <a href="https://docs.osmedeus.org/donation/"><img src="https://img.shields.io/badge/Sponsors-0078D4?style=for-the-badge&logo=GitHub-Sponsors&logoColor=39ff14&labelColor=black&color=black"></a>
  <a href="https://twitter.com/OsmedeusEngine"><img src="https://img.shields.io/badge/%40OsmedeusEngine-0078D4?style=for-the-badge&logo=Twitter&logoColor=39ff14&labelColor=black&color=black"></a>
  <a href="https://discord.gg/gy4SWhpaPU"><img src="https://img.shields.io/badge/Discord%20Server-0078D4?style=for-the-badge&logo=Discord&logoColor=39ff14&labelColor=black&color=black"></a>
  <a href="https://discord.gg/gy4SWhpaPU"><img src="https://img.shields.io/github/release/j3ssie/osmedeus?style=for-the-badge&labelColor=black&color=2fc414&logo=Github"></a>
  </p>
</p>

***

## Installation

> NOTE that you need some essential tools like `curl, wget, git, zip` and login as **root** to start

```shell
bash -c "$(curl -fsSL https://raw.githubusercontent.com/osmedeus/osmedeus-base/master/install.sh)"
```

## Build the engine from source

Make sure you installed `golang >= v1.17`

```shell
mkdir -p $GOPATH/src/github.com/j3ssie
git clone --depth=1 https://github.com/j3ssie/osmedeus $GOPATH/src/github.com/j3ssie/osmedeus
cd $GOPATH/src/github.com/j3ssie/osmedeus
make build
```

## Usage

```shell
# Scan Usage:
  osmedeus scan -f [flowName] -t [target]
  osmedeus scan -m [modulePath] -T [targetsFile]
  osmedeus scan -f /path/to/flow.yaml -t [target]
  osmedeus scan -m /path/to/module.yaml -t [target] --params 'port=9200'
  osmedeus scan -m /path/to/module.yaml -t [target] -l /tmp/log.log
  cat targets | osmedeus scan -f sample

# Practical Scan Usage:
  osmedeus scan -T list_of_targets.txt -W custom_workspaces
  osmedeus scan -t target.com -w workspace_name --debug
  osmedeus scan -f general -t www.sample.com
  osmedeus scan -f gdirb -T list_of_target.txt
  osmedeus scan -m ~/.osmedeus/core/workflow/test/dirbscan.yaml -t list_of_urls.txt
  osmedeus scan --wfFolder ~/custom-workflow/ -f your-custom-workflow -t list_of_urls.txt

# Provider Usage:
  osmedeus provider build
  osmedeus provider build --token xxx --rebuild --ic
  osmedeus provider create --name 'sample'
  osmedeus provider health --debug

# Cloud Usage:
  osmedeus cloud -f [flowName] -t [target]
  osmedeus cloud -m [modulePath] -t [target]
  osmedeus cloud -c 10 -f [flowName] -T [targetsFile]
  osmedeus cloud --token xxx -G -c 10 -f [flowName] -T [targetsFile]
  osmedeus cloud --chunk -c 10 -f [flowName] -t [targetsFile]

# Utilities Usage:
  osmedeus health
  osmedeus version --json
  osmedeus utils tmux ls
  osmedeus utils tmux logs -A -l 10
  osmedeus utils ps
  osmedeus utils ps --proc 'jaeles'
  osmedeus utils cron --cmd 'osmdeus scan -t example.com' --sch 60
  osmedeus utils cron --for --cmd 'osmedeus scan -t example.com'
```

## 💬 Community & Discussion

Join Our Discord server [here](https://discord.gg/gy4SWhpaPU)

## 💎 Donation & Sponsor

<p align="center">
 <img alt="Osmedeus" src="https://raw.githubusercontent.com/osmedeus/assets/main/premium-package.gif" />

 <p align="center"> Check out for a couple of <strong><a href="https://docs.osmedeus.org/donation/">donation methods here</a></strong> to get a <strong><a href="https://docs.osmedeus.org/premium/">premium package</a></strong><p>
</p>

## License

`Osmedeus` is made with ♥ by [@j3ssiejjj](https://twitter.com/j3ssiejjj) and it is released under the MIT license.
