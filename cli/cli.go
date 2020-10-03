/*
 * Copyright (c) 2020 Alex <alex@webz.asia>
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *   @author Alex <alex@webz.asia>
 *   @copyright Copyright (c) 2020 Alex <alex@webz.asia>
 *   @since 01.10.2020
 *
 */

package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/gh-replay/replay-go/cli/command"
)

const (
	UsageCommonFmt       = "There is a list of commands available: %v\n"
	CommandReplayCommits = "replay-commits"
)

type Args struct {
	Command       *string
	UsageCommon   func()
	cmds          map[string]*flag.FlagSet
	cmdList       []string
	ReplayCommits *command.ReplayCommitsArgs
}

// New instantiates Args with some init steps
func New() *Args {
	args := &Args{
		cmds: make(map[string]*flag.FlagSet),
	}

	args.newReplayCommitsCmd()

	args.cmdList = []string{}
	for k := range args.cmds {
		args.cmdList = append(args.cmdList, k)
	}
	sort.Strings(args.cmdList)
	args.UsageCommon = func() {
		fmt.Fprintf(flag.CommandLine.Output(), UsageCommonFmt, args.cmdList)
	}

	return args
}

// Usage prints all flag sets usages
func (a *Args) Usage() {
	flag.Usage()
	a.UsageCommon()
	for _, k := range a.cmdList {
		a.cmds[k].Usage()
	}
}

func (a *Args) newReplayCommitsCmd() {
	a.cmds[CommandReplayCommits] = flag.NewFlagSet(CommandReplayCommits, flag.ExitOnError)
	a.ReplayCommits = &command.ReplayCommitsArgs{
		InitStartDateString: a.cmds[CommandReplayCommits].String("new-init-date", time.Now().String(), "new initial commit starting date"),
		NewAuthor:           a.cmds[CommandReplayCommits].String("author", "", "new author name"),
	}
}

// Cli parses the commandline args and returns appropriate structure and/or error
func Cli() (args *Args, err error) {
	args = New()

	if len(os.Args) < 2 {
		err = errors.New("expected a command")

		return args, err
	}

	switch os.Args[1] {

	case CommandReplayCommits:
		args.cmds[CommandReplayCommits].Parse(os.Args[2:])
		err = args.ReplayCommits.Parse()

	default:
		err = errors.New("expected a command")
	}

	if err == nil {
		args.Command = &os.Args[1]
	}

	return args, err
}
