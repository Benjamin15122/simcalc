package simcalc

import (
	"bytes"
	"fmt"
	"math"
	"os/exec"

	maxflow "github.com/daviddengcn/go-algs/maxflow"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	lcs "github.com/yudai/golcs"
)

// LCSComputer calculates the longest common substring of two strings
func LCSComputer(lefts, rights string) (int, []interface{}) {
	leftBytes := []byte(lefts)
	left := make([]interface{}, len(leftBytes))
	for i, v := range leftBytes {
		left[i] = v
	}

	rightBytes := []byte(rights)
	right := make([]interface{}, len(rightBytes))
	for i, v := range rightBytes {
		right[i] = v
	}

	result := lcs.New(left, right)
	return result.Length(), result.Values()

}

// IRGenerator converts a code file into a llvm ir file
func IRGenerator(code, ll string) error {
	cmd := exec.Command("clang", "-S", "-emit-llvm", "-o", ll, code)
	err := cmd.Start()
	return err
}

// CleanUp do the tempfile removal job
func CleanUp() error {
	cmd := exec.Command("rm", "wgcltemp", "wgcrtemp")
	err := cmd.Start()
	return err
}

// FSGenerator converts a ll file into a string[] where each string represents a function
func FSGenerator(ll string) ([]string, error) {
	m, err := asm.ParseFile(ll)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	var funclist []string
	// For each function of the module.
	for _, f := range m.Funcs {
		// For each basic block of the function.
		for _, block := range f.Blocks {
			funcs := ""
			// For each non-branching instruction of the basic block.
			for _, inst := range block.Insts {
				// Type switch on instruction to find call instructions.
				switch inst := inst.(type) {
				// case *ir.InstCall:
				// 	callee := inst.Callee.Ident()
				// 	// Add edges from caller to callee.
				// 	fmt.Fprintf(buf, "\t%q -> %q\n", caller, callee)
				default:
					fmt.Fprintf(buf, "%T", inst)
					insts := buf.String()
					buf.Reset()
					len := len(insts)
					if len < 12 {
						funcs = funcs + insts[8:len]
					} else {
						funcs = funcs + insts[8:12]
					}
				}
			}
			funclist = append(funclist, funcs)
			// Terminator of basic block.
			switch term := block.Term.(type) {
			case *ir.TermRet:
				// do something.
				_ = term
			}
		}
	}
	return funclist, nil
}

// MaxFlow calculates the maxflow of p1 and p2
func MaxFlow(p1, p2 []string, mode string) (float32, error) {
	g := maxflow.NewGraph()
	len1 := len(p1)
	len2 := len(p2)
	maxlen := len1 + len2 + 1
	nodes := make([]*maxflow.Node, len1+len2+2)
	for i := range nodes {
		nodes[i] = g.AddNode()
	}
	//set source
	g.SetTweights(nodes[0], 0, maxflow.CapType(math.MaxInt32))
	//set sink
	g.SetTweights(nodes[maxlen], maxflow.CapType(math.MaxInt32), 0)
	//set rest nodes
	for i := 1; i < maxlen; i++ {
		g.SetTweights(nodes[i], maxflow.CapType(math.MaxInt32), maxflow.CapType(math.MaxInt32))
	}
	//set source edges to p1
	for i := 1; i <= len1; i++ {
		g.AddEdge(nodes[0], nodes[i], maxflow.CapType(len(p1[i-1])), 0)
	}
	//set p2 to sink edges
	for i := len1 + 1; i <= len1+len2; i++ {
		g.AddEdge(nodes[i], nodes[maxlen], maxflow.CapType(len(p2[i-len1-1])), 0)
	}
	//set rest edges
	for i, fi := range p1 {
		for j, fj := range p2 {
			c := 0
			switch mode {
			case "LCS":
				c, _ = LCSComputer(fi, fj)
			default:
				if len(fi) > len(fj) {
					c = len(fi)
				} else {
					c = len(fj)
				}
			}
			g.AddEdge(nodes[i+1], nodes[j+len1], maxflow.CapType(c), 0)
		}
	}

	g.Run()

	flow := g.Flow()
	return float32(flow), nil
}

// SimCalculate calculates similarity prints a float64 value ranging [0-100].
func SimCalculate(code1, code2 string) error {
	//compile code to llvm ir
	err := IRGenerator(code1, "wgcltemp")
	if err != nil {
		cerr := CleanUp()
		if cerr != nil {
			return cerr
		}
		return err
	}
	err = IRGenerator(code2, "wgcrtemp")
	if err != nil {
		//clean up
		cerr := CleanUp()
		if cerr != nil {
			return cerr
		}
		return err
	}
	//generate function list p1 and p2
	p1, gerr1 := FSGenerator("wgcltemp")
	if gerr1 != nil {
		//clean up
		cerr := CleanUp()
		if cerr != nil {
			return cerr
		}
		return gerr1
	}
	p2, gerr2 := FSGenerator("wgcrtemp")
	if gerr2 != nil {
		//clean up
		cerr := CleanUp()
		if cerr != nil {
			return cerr
		}
		return gerr2
	}
	//maxflow
	// lcsf, _ := MaxFlow(p1, p2, "LCS")
	// mf, _ := MaxFlow(p1, p2, "")
	// fmt.Print(lcsf,mf)
	var pr []float32
	ttl := 0
	for _, fi := range p1 {
		var pri float32
		pri = 0
		for _, fj := range p2 {
			c, _ := LCSComputer(fi, fj)
			prt := float32(c) / float32(len(fi))
			if prt > pri {
				pri = prt
			}
		}
		ttl = ttl + len(fi)
		pr = append(pr, pri)
	}
	var p float32
	p = 0
	for i, fi := range p1 {
		p = p + pr[i]*float32(len(fi))/float32(ttl)
	}
	fmt.Printf("%.2f%%\n", p*100)
	//cleanup
	cerr := CleanUp()
	if cerr != nil {
		return cerr
	}
	return nil
}

// DebugCalculate calculates similarity and prints process log.
func DebugCalculate(code1, code2 string) error {
	fmt.Println("It's a checking process with verbosed option and would print debugging information")
	fmt.Println("Path of code1: " + code1)
	fmt.Println("Path of code2: " + code2)
	//compile code to llvm ir
	err := IRGenerator(code1, "wgcltemp")
	if err != nil {
		fmt.Println("Something went wrong when generating ir files of " + code1)
		//clean up
		cerr := CleanUp()
		if cerr != nil {
			fmt.Println("Something went wrong when cleaning up, please delete wgcltemp wgcrtemp manually if they are still there")
		}
		return err
	}
	err = IRGenerator(code2, "wgcrtemp")
	if err != nil {
		fmt.Println("Something went wrong when generating ir files of " + code2)
		//clean up
		cerr := CleanUp()
		if cerr != nil {
			fmt.Println("Something went wrong when cleaning up, please delete wgcltemp wgcrtemp manually if they are still there")
		}
		return err
	}
	//generate function list p1 and p2
	p1, gerr1 := FSGenerator("wgcltemp")
	if gerr1 != nil {
		fmt.Println("Something went wrong when generating functionlist and machine inst of " + code1)
		//clean up
		cerr := CleanUp()
		if cerr != nil {
			fmt.Println("Something went wrong when cleaning up, please delete wgcltemp wgcrtemp manually if they are still there")
		}
		return gerr1
	}
	p2, gerr2 := FSGenerator("wgcrtemp")
	if gerr2 != nil {
		fmt.Println("Something went wrong when generating functionlist and machine inst of " + code2)
		//clean up
		cerr := CleanUp()
		if cerr != nil {
			fmt.Println("Something went wrong when cleaning up, please delete wgcltemp wgcrtemp manually if they are still there")
		}
		return gerr2
	}
	//maxflow
	// lcsf, _ := MaxFlow(p1, p2, "LCS")
	// mf, _ := MaxFlow(p1, p2, "")
	// fmt.Print(lcsf,mf)
	var pr []float32
	ttl := 0
	for _, fi := range p1 {
		var pri float32
		pri = 0
		for _, fj := range p2 {
			c, _ := LCSComputer(fi, fj)
			prt := float32(c) / float32(len(fi))
			if prt > pri {
				pri = prt
			}
		}
		ttl = ttl + len(fi)
		pr = append(pr, pri)
	}
	var p float32
	p = 0
	for i, fi := range p1 {
		p = p + pr[i]*float32(len(fi))/float32(ttl)
	}
	fmt.Printf("%.2f%%\n", p*100)
	//cleanup
	cerr := CleanUp()
	if cerr != nil {
		fmt.Println("Something went wrong when cleaning up, please delete wgcltemp wgcrtemp manually if they are still there")
	}
	return nil
}
