package main

import (
    sp "github.com/scipipe/scipipe"
)

func main() {
    wf := sp.NewWorkflow("pipeline", 1)

    wf.NewProc("fastqc", "bash -c 'fastqc /home/alowator/bio/pipeline/SRR24631089_1.fastq > res.out'")

    wf.Run()
}
