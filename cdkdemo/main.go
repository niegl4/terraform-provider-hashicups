package main

import (
	order2 "cdk.tf/go/stack/generated/edu/hashicups/order"
	"cdk.tf/go/stack/generated/edu/hashicups/provider"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here
	host := "http://localhost:19090"
	password := "education"
	username := "test123"
	provider.NewHashicupsProvider(stack, jsii.String("provider_demo"), &provider.HashicupsProviderConfig{
		Alias:    nil,
		Host:     &host,
		Password: &password,
		Username: &username,
	})

	coffeeId := 123.456
	quantityId := 123.456
	order2.NewOrder(stack, jsii.String("coffee_demo"), &order2.OrderConfig{
		Connection:   nil,
		Count:        nil,
		DependsOn:    nil,
		ForEach:      nil,
		Lifecycle:    nil,
		Provider:     nil,
		Provisioners: nil,
		Items: []order2.OrderItems{
			{
				Coffee: &order2.OrderItemsCoffee{
					Id: &coffeeId,
				},
				Quantity: &quantityId,
			},
		},
	})

	return stack
}

/*
实验环境：
1.安装node.js和npm，官网一个package全都包括了，设置环境变量。
2.npm指向国内镜像源：
npm config set registry https://registry.npmmirror.com
npm config get registry
3.npm install --global cdktf-cli@latest
4.我的本地启动了插件服务，并设置了reattach环境变量
export TF_REATTACH_PROVIDERS='{"hashicorp.com/edu/hashicups":{"Protocol":"grpc","ProtocolVersion":6,"Pid":6069,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/9s/81g75_zn7sbbwfswjpq352bh0000gn/T/plugin3441574700"}}
5.设置cdktf的环境变量与配置文件，不过好像由于上一步的存在，这一步似乎没起作用
export TF_CLI_CONFIG_FILE=/Users/nieguanglin/go/src/github.com/niegl4/terraform-provider-hashicups/cdkconfig/dev.tfrc
dev.tfrc中有如下内容：
provider_installation {

	  # Use /home/developer/tmp/terraform-null as an overridden package directory
	  # for the hashicorp/null provider. This disables the version and checksum
	  # verifications for this provider and forces Terraform to look for the
	  # null provider plugin in the given directory.
	  dev_overrides {
	    "hashicorp.com/edu/hashicups" = "/Users/nieguanglin/go/bin/terraform-provider-hashicups"
	  }

	  # For all other providers, install them directly from their origin provider
	  # registries as normal. If you omit this, Terraform will _only_ use
	  # the dev_overrides block, and so no other providers will be available.
	  direct {}
	}

6.初始化项目。虽然提前做了配置，但这一步仍然报错，不过基本框架已经生成。
cdktf init --template=go --providers=hashicorp.com/edu/hashicups --local
7.修改.cdktf.json文件

	{
	  "language": "go",
	  "app": "go run main.go",
	  "codeMakerOutput": "generated",
	  "projectId": "26d1c3bc-4ad5-446f-a624-b9f2b34a69c8",
	  "sendCrashReports": "false",
	  "terraformProviders": [
	    "hashicorp.com/edu/hashicups"
	  ],
	  "terraformModules": [],
	  "context": {

	  }
	}

8.自动生成代码
cdktf get
9.写go代码，通过cdktf生成json模板
cdktf synth --json
*/
func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdkdemo")

	app.Synth()
}
