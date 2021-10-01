//+build mage

package main

import "os"

func Build() error {
	return nil
}

func Clean() error {
	if err := os.Remove("ff-card-be.exe"); err != nil {
		return err
	}

	return os.Remove("conf.json")
}

func Install() error {
	return nil
}

func Uninstall() error {
	return nil
}
