// Package model contains the structure and functions used to pass information
// between the computations of the bot and the driver code, in bot.go
package model

// CommandResponse is returned by the command handler, it contains information
// about the command computation.
type CommandResponse struct {
	Text        string
	NextCommand string
}

// makeResponse creates a CommandResponse with the given text and nextCommand
func makeResponse(text string, nextCommand string) CommandResponse {
	return CommandResponse{
		Text:        text,
		NextCommand: nextCommand,
	}
}

// makeResponseWithText creates a CommandResponse with the given text (and no nextCommand)
func makeResponseWithText(text string) CommandResponse {
	return makeResponse(text, "")
}

// makeResponseWithNextCommand creates a CommandResponse with the given nextCommand (and no text)
func makeResponseWithNextCommand(nextCommand string) CommandResponse {
	return makeResponse("", nextCommand)
}

// IsEmpty returns true if the CommandResponse has no text and no nextCommand
func (r CommandResponse) IsEmpty() bool {
	return r.Text == "" && r.NextCommand == ""
}

// HasText returns true if the CommandResponse has some text
func (r CommandResponse) HasText() bool {
	return r.Text != ""
}

// HasNextCommand returns true if the CommandResponse has some nextCommand
func (r CommandResponse) HasNextCommand() bool {
	return r.NextCommand != ""
}
