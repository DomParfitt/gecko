export interface IGeckoAutomata {
    CurrentState: number,
    TerminalStates: number[],
    States: number[],
	Transitions: any // map[int]map[string]int
}