#!/bin/bash
antlr -Dlanguage=Go -package antlr -visitor ./parser/antlr/*.g4 -o ./parser/antlr