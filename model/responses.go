// This file contains the structure and functions used to pass information
// between the computations of the bot and the driver code, in bot.go
package model

// This struct is returned by the command handler, it contains information about the
// command computation.
type CommandResponse struct {
	Text        string
	NextCommand string
}

func makeResponse(text string, nextCommand string) CommandResponse {
	return CommandResponse{
		Text:        text,
		NextCommand: nextCommand,
	}
}

func makeResponseWithText(text string) CommandResponse {
	return makeResponse(text, "")
}

func makeResponseWithNextCommand(nextCommand string) CommandResponse {
	return makeResponse("", nextCommand)
}

func (respose CommandResponse) IsEmpty() bool {
	return respose.Text == "" && respose.NextCommand == ""
}

func (response CommandResponse) HasText() bool {
	return response.Text != ""
}

func (response CommandResponse) HasNextCommand() bool {
	return response.NextCommand != ""
}
