package cmd

import (
    "fmt"
    "github.com/fatih/color"
    "github.com/j3ssie/osmedeus/core"
    "github.com/j3ssie/osmedeus/libs"
    "github.com/spf13/cobra"
)

// RootUsage base help
func RootUsage() {
    var h string
    h += ScanUsage()
    h += CloudUsage()
    h += UtilsUsage()

    fmt.Println(h)
}

func ScanUsage() string {
    h := color.HiCyanString("Example Scan Commands:")
    h += `
  ## Start a simple scan with default 'general' flow
  osmedeus scan -t sample.com
  
  ## Start a general scan but exclude some of the module
  osmedeus scan -t sample.com -x screenshot -x spider
  
  ## Start a simple scan with other flow
  osmedeus scan -f vuln -t sample.com
  
  ## Scan for CIDR with file contains CIDR with the format '1.2.3.4/24'
  osmedeus scan -f cidr -t list-of-cidrs.txt
  osmedeus scan -f cidr -t '1.2.3.4/24' # this will auto convert the single input to the file and run
  
  ## Directly run on vuln scan and directory scan on list of domains
  osmedeus scan -f vuln-and-dirb -t list-of-domains.txt
  
  ## Use a custom wordlist
  osmedeus scan -t sample.com -p 'wordlists={{.Data}}/wordlists/content/big.txt' -p 'fthreads=40'
  
  ## Scan list of targets
  osmedeus scan -T list_of_targets.txt
  
  ## Get target from a stdin and start the scan with 2 concurrency
  cat list_of_targets.txt | osmedeus scan -c 2
            `
    h += color.HiCyanString("\nScan Usage:\n")
    h += "  osmedeus scan -f [flowName] -t [target] \n"
    h += "  osmedeus scan -m [modulePath] -T [targetsFile] \n"
    h += "  osmedeus scan -f /path/to/flow.yaml -t [target] \n"
    h += "  osmedeus scan -m /path/to/module.yaml -t [target] --params 'port=9200'\n"
    h += "  osmedeus scan -m /path/to/module.yaml -t [target] -l /tmp/log.log\n"
    h += "  cat targets | osmedeus scan -f sample\n"

    h += color.HiCyanString("\nPractical Scan Usage:\n")
    h += "  osmedeus scan -T list_of_targets.txt -W custom_workspaces\n"
    h += "  osmedeus scan -t target.com -w workspace_name --debug\n"
    h += "  osmedeus scan -f general -t www.sample.com\n"
    h += "  osmedeus scan -f gdirb -T list_of_target.txt\n"
    h += "  osmedeus scan -m ~/.osmedeus/core/workflow/test/dirbscan.yaml -t list_of_urls.txt\n"
    return h
}

func UtilsUsage() string {
    h := color.HiCyanString("\nUtilities Usage:\n")
    h += "  osmedeus health \n"
    h += "  osmedeus version --json \n"
    h += "  osmedeus utils tmux ls \n"
    h += "  osmedeus utils tmux logs -A -l 10 \n"
    h += "  osmedeus utils ps \n"
    h += "  osmedeus utils ps --proc 'jaeles' \n"
    h += "  osmedeus utils cron --cmd 'osmdeus scan -t example.com' --sch 60\n"
    h += "  osmedeus utils cron --for --cmd 'osmedeus scan -t example.com'\n"
    return h
}

func ConfigUsage() string {
    h := color.HiCyanString("\nConfig Usage:\n")
    h += "  osmedeus config [action] [OPTIONS] \n"
    h += "  osmedeus config init -p https://github.com/j3ssie/osmedeus-plugins\n"
    h += "  osmedeus config --user newusser --pass newpassword\n"
    h += "  osmedeus config reload \n"
    h += "  osmedeus config update \n"
    h += "  osmedeus config clean \n"
    h += "  osmedeus config delete -t woskapce \n"
    h += "  osmedeus config delete -w workspace_name \n"
    return h
}

func CloudUsage() string {
    h := color.HiCyanString("\nProvider Usage:\n")
    h += "  osmedeus provider build \n"
    h += "  osmedeus provider build --token xxx --rebuild --ic\n"
    h += "  osmedeus provider create --name 'sample' \n"
    h += "  osmedeus provider health --debug \n"

    h += color.HiCyanString("\nCloud Usage:\n")
    h += "  osmedeus cloud -f [flowName] -t [target] \n"
    h += "  osmedeus cloud -m [modulePath] -t [target] \n"
    h += "  osmedeus cloud -c 10 -f [flowName] -T [targetsFile] \n"
    h += "  osmedeus cloud --token xxx -G -c 10 -f [flowName] -T [targetsFile] \n"
    h += "  osmedeus cloud --chunk -c 10 -f [flowName] -t [targetsFile] \n"

    return h
}

// ScanHelp scan help message
func ScanHelp(cmd *cobra.Command, _ []string) {
    fmt.Println(core.Banner())
    fmt.Println(cmd.UsageString())
    h := ScanUsage()
    fmt.Println(h)
    printDocs()
}

// CloudHelp scan help message
func CloudHelp(cmd *cobra.Command, _ []string) {
    fmt.Println(core.Banner())
    fmt.Println(cmd.UsageString())
    h := CloudUsage()
    fmt.Println(h)
    printDocs()
}

// ConfigHelp config help message
func ConfigHelp(cmd *cobra.Command, _ []string) {
    fmt.Println(core.Banner())
    fmt.Println(cmd.UsageString())
    h := ConfigUsage()
    fmt.Println(h)
    printDocs()
}

// UtilsHelp utils help message
func UtilsHelp(cmd *cobra.Command, _ []string) {
    fmt.Println(core.Banner())
    fmt.Println(cmd.UsageString())
    h := UtilsUsage()
    fmt.Println(h)
    printDocs()
}

// RootHelp print help message
func RootHelp(cmd *cobra.Command, _ []string) {
    fmt.Println(core.Banner())
    fmt.Println(cmd.UsageString())
    RootUsage()
    printDocs()
}

func printDocs() {
    fmt.Printf("Documentation can be found here: %s\n", color.GreenString(libs.DOCS))
}
