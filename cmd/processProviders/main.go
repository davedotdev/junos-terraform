// Copyright (c) 2017-2021, Juniper Networks Inc. All rights reserved.
//
// License: Apache 2.0
//
// THIS SOFTWARE IS PROVIDED BY Juniper Networks, Inc. ''AS IS'' AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL Juniper Networks, Inc. BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Juniper/junos-terraform/Internal/cfg"
	"github.com/Juniper/junos-terraform/Internal/processProviders"
)

// Syntactic helper to reduce repetition.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// The user can pass the configuration as part of config file or command line arguments.
	jcfg := cfg.Config{}

	// Get the config location
	flagConfig := flag.String("config", "", "Path to config file")

	// Get flags from the user for JTAF
	flagYang := flag.String("yangDir", "", "Absolute path to Yang files directory")
	flagXpath := flag.String("xpathPath", "", "Absolute path to file with xpath for Providers")
	flagProvider := flag.String("providerDir", "", "Absolute path to directory to generate provider")

	// Provide a line break for pretty reasons.
	fmt.Println()

	flag.Parse()

	// Get the config.
	if *flagConfig != "" {
		jcfg, err := cfg.GetConfig(*flagConfig)
		if err != nil {
			fmt.Println("Error retrieving configuration: ", err)
		}
		check(processProviders.CreateProviders(jcfg))
	} else if *flagYang != "" || *flagXpath != "" || *flagProvider != "" {
		// If config file path is not present then check for individual elements.
		jcfg.XpathPath = *flagXpath
		jcfg.ProviderDir = *flagProvider
		jcfg.YangDir = *flagYang

		check(processProviders.CreateProviders(jcfg))
	} else {
		fmt.Println("One or more mandatory inputs are missing, exiting...")
		os.Exit(0)
	}
}
