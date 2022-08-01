EVM Smart Contracts Similarity
==============================

Detect identical and similar smart contracts based on their bytecode.

Requirements
------------

* An installation of golang

```
go mod tidy
```

Usage
-----

```
go run src/cmd/analyze/main.go --help
Usage: main [--json JSON] [--debug] [--skipproxy] [--ngram NGRAM] [--threshold THRESHOLD] PATH

Positional arguments:
  PATH                   Path to directory containing bytecodes

Options:
  --json JSON            Filepath to save the output as a JSON
  --debug                Print debug information
  --skipproxy            Skip potential proxy contracts
  --ngram NGRAM          Select how many elements should an n-gram have [default: 5]
  --threshold THRESHOLD
                         Set a similarity threshold [default: 90]
  --help, -h             display this help and exit
```

Detect Identical and Similar Contracts
--------------------------------------

First, it will search for files ending in `.bin` in the directory 
(and the sub-directory) provided by the CLI.
Then, it will detect clusters for the following:

1. Identical contracts: The bytecode is the same.
2. Opcode-identical contracts: After removing the push arguments, 
the disassembled opcodes are identical.
3. Similar contracts: compute n-grams of Ethereum opcodes and use Jaccard similarity to cluster contracts. Note that a contract belongs to the first cluster it matches based on the given threshold

If you save the output to a JSON file, the outcome would be similar to the following.

```
{
  "identical": [
    [
      "0x0624eb9691d99178d0d2bd76c72f1dbb4db05286"
    ],
    [
      "0x0e511aa1a137aad267dfe3a6bfca0b856c1a3682"
    ],
    [
      "0x32e574b0400cdb93dc33a34a30986737186bde43",
      "0xfe670043c158adc057b5ef2f66e753847e1c3d1d"
    ]
  ],
  "opcodes": [
    [
      "0x0624eb9691d99178d0d2bd76c72f1dbb4db05286"
    ],
    [
      "0x0e511aa1a137aad267dfe3a6bfca0b856c1a3682"
    ],
    [
      "0x32e574b0400cdb93dc33a34a30986737186bde43",
      "0xfe670043c158adc057b5ef2f66e753847e1c3d1d"
    ]
  ],
  "similarity": [
    [
      "0x0624eb9691d99178d0d2bd76c72f1dbb4db05286"
    ],
    [
      "0x0e511aa1a137aad267dfe3a6bfca0b856c1a3682",
      "0x32e574b0400cdb93dc33a34a30986737186bde43",
      "0xfe670043c158adc057b5ef2f66e753847e1c3d1d"
    ]
  ]
}
```

Relevant Papers
---------------

* Detect similar contracts using n-grams and cosine similarity. [Analyzing Ethereum’s Contract Topology](https://www.ccs.neu.edu/home/amislove/publications/Ethereum-IMC.pdf)
* Detect similar contracts after removing creation code and swarm code, while using fuzzy hashing and pair-wise similarity. [Characterizing Code Clones in the Ethereum Smart Contract Ecosystem](https://fc20.ifca.ai/preproceedings/106.pdf)
* Use smart contract birthmarks to enable clone detection. [Enabling clone detection for ethereum via smart contract birthmarks](https://ieeexplore.ieee.org/document/8813297)
* Clone detection based on CFG. [Similarity Measure for Smart Contract Bytecode Based on CFG Feature Extraction](https://ieeexplore.ieee.org/abstract/document/9718856)
