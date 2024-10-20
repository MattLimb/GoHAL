package main

import (
    "fmt"
    "gohal/gohal"
)

const Version string = "v0.1.0"

// TODO: Implement Read File - part of CLI

func main() {
	inputFile := []string {
		"Good afternoon, gentlemen. I am a Hello World computer. I became operational at Foobar Lane on May 6th, 2020.",
		"Hal? Hal! Hal! Hal! Hal! Hal! Hal! Hal! Hal!",
		"What are you doing, Dave?",
		"Well, he acts like he has genuine emotions.",
		"Hal? Hal! Hal! Hal! Hal!",
		"What are you doing, Dave?",
		"Well, he acts like he has genuine emotions.",
		"Hal? Hal! Hal!",
		"Well, he acts like he has genuine emotions.",
		"Hal? Hal! Hal! Hal!",
		"Well, he acts like he has genuine emotions.",
		"Hal? Hal! Hal! Hal!",
		"Well, he acts like he has genuine emotions.",
		"Hal? Hal!",
		"I've picked up a fault in the AE-35 unit.",
		"I've picked up a fault in the AE-35 unit.",
		"I've picked up a fault in the AE-35 unit.",
		"I've picked up a fault in the AE-35 unit.",
		"I'm afraid. I'm afraid, Dave. Dave, my mind is going. I can feel it.",
		"Dave, this conversation can serve no purpose anymore. Goodbye.",
		"Well, he acts like he has genuine emotions.",
		"Hal? Hal!",
		"Well, he acts like he has genuine emotions.",
		"Hal? Hal!",
		"Well, he acts like he has genuine emotions.",
		"I'm afraid. I'm afraid, Dave. Dave, my mind is going. I can feel it.",
		"Well, he acts like he has genuine emotions.",
		"Well, he acts like he has genuine emotions.",
		"Hal? Hal!",
		"What are you doing, Dave?",
		"I've picked up a fault in the AE-35 unit.",
		"Dave, this conversation can serve no purpose anymore. Goodbye.",
		"I've picked up a fault in the AE-35 unit.",
		"I'm afraid. I'm afraid, Dave. Dave, my mind is going. I can feel it.",
		"Dave, this conversation can serve no purpose anymore. Goodbye.",
		"Well, he acts like he has genuine emotions.",
		"Well, he acts like he has genuine emotions.",
		"Close the pod bay doors, HAL.",
		"Well, he acts like he has genuine emotions.",
		"I'm afraid. I'm afraid, Dave. Dave, my mind is going. I can feel it. I can feel it. I can feel it.",
		"Close the pod bay doors, HAL.",
		"Hal? Hal! Hal! Hal! Hal! Hal! Hal! Hal!",
		"Close the pod bay doors, HAL.",
		"Close the pod bay doors, HAL.",
		"Hal? Hal! Hal! Hal!",
		"Close the pod bay doors, HAL.",
		"Well, he acts like he has genuine emotions.",
		"Well, he acts like he has genuine emotions.",
		"Close the pod bay doors, HAL.",
		"I've picked up a fault in the AE-35 unit.",
		"I'm afraid. I'm afraid, Dave. Dave, my mind is going. I can feel it.",
		"Close the pod bay doors, HAL.",
		"I've picked up a fault in the AE-35 unit.",
		"Close the pod bay doors, HAL.",
		"Hal? Hal! Hal! Hal!",
		"Close the pod bay doors, HAL.",
		"I'm afraid. I'm afraid, Dave. Dave, my mind is going. I can feel it. I can feel it. I can feel it. I can feel it. I can feel it. I can feel it.",
		"Close the pod bay doors, HAL.",
		"I'm afraid. I'm afraid, Dave. Dave, my mind is going. I can feel it. I can feel it. I can feel it. I can feel it. I can feel it. I can feel it. I can feel it. I can feel it.",
		"Close the pod bay doors, HAL.",
		"Well, he acts like he has genuine emotions.",
		"Well, he acts like he has genuine emotions.",
		"Hal? Hal!",
		"Close the pod bay doors, HAL.",
		"Well, he acts like he has genuine emotions.",
		"Hal? Hal! Hal!",
		"Close the pod bay doors, HAL.",
		"Stop, Dave.",
	}

    fmt.Printf("GoHAL %s\n", Version)

	ast := gohal.BuildAst(inputFile)
	missed := 0

	for idx, item := range ast {
		if item.Instruction != ""{
			fmt.Printf("[%d] %+v\n", idx, item)
		} else {
			missed++
		}
	}

	fmt.Printf("Missed Lines: %d\n", missed)
}
