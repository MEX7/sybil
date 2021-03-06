package sybil_cmd

import sybil "github.com/logv/sybil/src/lib"

import "flag"

// our aggregate command will take multiple output files on the command in,
// verify all their query specs match the same md5 result and then combine them

func RunAggregateCmdLine() {
	flag.Parse()
	dirs := flag.Args()

	var t, f = true, false
	sybil.FLAGS.DEBUG = &t
	sybil.Debug("AGGREGATING")

	sybil.DecodeFlags()
	sybil.FLAGS.PRINT = &t
	sybil.FLAGS.ENCODE_RESULTS = &f
	sybil.Debug("AGGREGATING DIRS", dirs)

	vt := sybil.VTable{}
	vt.StitchResults(dirs)

}
