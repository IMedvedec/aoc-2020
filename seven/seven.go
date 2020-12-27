package seven

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	filename = "input/seven_input"
)

// bagType as type string could be optimised.
type bagType string

// node represents a graph data structure entity.
type node struct {
	bag bagType
	// containers represents a set of bags in which the node can be found.
	containers map[bagType]*node
}

// newNode is a node constructor.
func newNode(bag bagType) node {
	newNode := node{
		bag:        bag,
		containers: make(map[bagType]*node),
	}

	return newNode
}

// addReferenceToOuterBagNode appends the outerBagNode into the node
// containers set.
func (n *node) addReferenceToOuterBagNode(outerBagNode *node) {
	if _, ok := n.containers[outerBagNode.bag]; !ok {
		n.containers[outerBagNode.bag] = outerBagNode
	}
}

// structure represents the bag graph structure.
type structure map[bagType]*node

// newStructure is a structure constructor.
func newStructure() structure {
	newStructure := make(map[bagType]*node)
	return newStructure
}

// possibleOuterBagCountForBag contains the problem solution.
func (s structure) possibleOuterBagCountForBag(bag bagType) int {
	visitedOuterNodes := make(map[bagType]struct{})

	var recursion func(node *node) int
	recursion = func(node *node) int {
		if _, ok := visitedOuterNodes[node.bag]; ok {
			return 0
		}
		visitedOuterNodes[node.bag] = struct{}{}

		var possibleOuterBagCount int
		for _, outerNode := range node.containers {
			possibleOuterBagCount += recursion(outerNode)
		}

		return possibleOuterBagCount + 1
	}

	// -1 reduces the starting node incremented return.
	return recursion(s[bag]) - 1
}

// rule defines a input rule.
type rule string

// getOuterBag parses the outer bag from the rule.
func (r rule) getOuterBag() bagType {
	innerAndOuterRuleParts := strings.Split(string(r), "contain")

	outerBagRule := strings.TrimSpace(innerAndOuterRuleParts[0])

	words := strings.Fields(outerBagRule)

	return bagType(strings.Join(words[:len(words)-1], " "))
}

// getInnerBags parses the inner bags from the rule.
func (r rule) getInnerBags() []bagType {
	innerAndOuterRuleParts := strings.Split(string(r), "contain")

	innerBagRule := innerAndOuterRuleParts[1]

	var bags []bagType

	innerBags := strings.Split(innerBagRule, ",")

	for _, innerBag := range innerBags {
		words := strings.Fields(strings.TrimSpace(innerBag))
		bags = append(bags, bagType(strings.Join(words[1:len(words)-1], " ")))
	}

	return bags
}

// Run is the solution starting point.
func Run() {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error while file opening: %v", err)
	}

	scanner := bufio.NewScanner(file)
	structure := newStructure()

	for scanner.Scan() {
		line := scanner.Text()
		analyseRule(structure, rule(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	log.Printf("Bag color count: %d\n", structure.possibleOuterBagCountForBag(bagType("shiny gold")))
}

// analyseRule is a helper function which processes a single input rule.
func analyseRule(structure structure, rule rule) {

	outerBag := rule.getOuterBag()
	innerBags := rule.getInnerBags()

	outerBagNode, ok := structure[outerBag]
	if !ok {
		newOuterBag := newNode(outerBag)
		structure[outerBag] = &newOuterBag
		outerBagNode = &newOuterBag
	}

	for _, innerBag := range innerBags {
		innerBagNode, ok := structure[innerBag]
		if !ok {
			newInnerBag := newNode(innerBag)
			structure[innerBag] = &newInnerBag
			innerBagNode = &newInnerBag
		}

		innerBagNode.addReferenceToOuterBagNode(outerBagNode)
	}
}
