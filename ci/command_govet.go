// +build mage

/*
 * Copyright (C) 2018 The "MysteriumNetwork/go-openvpn" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
	"github.com/dvnetwork/go-openvpn/ci/util"
)

// Checks that the source is compliant with go vet
func GoVet() error {
	path := os.Getenv(util.MagePathOverrideEnvVar)
	if path == "" {
		path = "../..."
	}
	out, err := sh.Output("go", "vet", "-composites=false", path)
	fmt.Print(out)
	if err != nil {
		return err
	}
	fmt.Println("All files are compliant with go vet")
	return nil
}
