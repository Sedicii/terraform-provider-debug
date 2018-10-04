# Debug Terraform Provider

This Terraform provides allows to log text or data structures to a file for debug purposes. 

### Maintainers

This provider plugin is maintained by [Sedicii](https://sedicii.com/).

### Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

### Installation

```bash
curl https://raw.githubusercontent.com/Sedicii/terraform-provider-debug/master/scripts/install-debug-tf-pluging.sh | bash
```

### Usage

```hcl-terraform
provider "debug" {
  log_file = "<path to log file>" // Default to "/tmp/tf-debug-provider.log"
}

locals {
    data = {
        test = 1
        bla = "test"
    }
}

data "debug_log" "log" {
   line = "bla value is ${local.data.bla}" // type string, for log lines
   data = "${local.data}"  // type map, for data structures
   tag = "some_tag" // will appear in the log line. Default to "default"
}
```
