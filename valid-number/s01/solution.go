package s01

// https://leetcode.com/problems/valid-number/discuss/512431/Golang-DFA-clean-and-understandable-0ms-beats-100

type state int // the state is used to lookup what next states are valid

const (
    invalidState state = iota // Meaningful zero values are idiomatic Go. With invalid being zero, our transform map is less verbose.
    initial // initial state, where we begin
    sign // + or - found on LHS of exponent
    digitNoDot // digit found, no dot yet found (LHS of exponent), eg. "1" of "1.2334"
    dotAndDigit // digit found and dot has already been found (LHS of exponent), eg. "1.2" of "1.2334"
    dotNoDigit // dot found, no digit yet found (LHS of exponent), eg. "." of ".0123
    exp // e found
    signPostExp // + or - found on RHS of exponent
    digitPostExp // digit found (RHS of exponent), no dots are allowed on the RHS so the digit handling is easier
    blank // trailing spaces
    done // EOF
)

type input int // inputs transform state

const (
    invalidInput input = iota // bad input
    digitInput // 0-9
    signInput // +/-
    expInput // e
    dotInput // .
    blankInput // space
    doneInput // eof
)

// Every state maintains a map of inputs to new states.
// If a state does not support an input, then Go will automatically return 0
// which is invalid state. This is useful because that is what we would
// expect if a state cannot take an input.
var transform = map[state]map[input]state{
    initial: {
        digitInput: digitNoDot,
        signInput: sign,
        dotInput: dotNoDigit,
        blankInput: initial,
    },
    sign: {
        digitInput: digitNoDot,
        dotInput: dotNoDigit,
    },
    digitNoDot: {
        digitInput: digitNoDot,
        dotInput: dotAndDigit,
        expInput: exp,
        blankInput: blank,
        doneInput: done,
    },
    dotAndDigit: {
        digitInput: dotAndDigit,
        expInput: exp,
        blankInput: blank,
        doneInput: done,
    },
    dotNoDigit: {
        digitInput: dotAndDigit,
    },
    exp: {
        signInput: signPostExp,
        digitInput: digitPostExp,
    },
    signPostExp: {
        digitInput: digitPostExp,
    },
    digitPostExp: {
        digitInput: digitPostExp,
        blankInput: blank,
        doneInput: done,
    },
    blank: {
        blankInput: blank,
        doneInput: done,
    },
	// done is missing because done has no states it can go to, so we leave it off
}

// Now that we have 99% of the code simply being declarative in nature,
// all we have to do is iterate over each character and try to transform the
// current state into the next state. After doing so, we make sure we are
// not in an invalid state and if we are, we can early exit with a failure.
func IsNumber(s string) bool {
    currentState := initial
    for _, chr := range s {
        switch chr {
        case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
            currentState = transform[currentState][digitInput]
        case '.':
            currentState = transform[currentState][dotInput]
        case '+', '-':
            currentState = transform[currentState][signInput]
        case 'e':
            currentState = transform[currentState][expInput]
        case ' ':
            currentState = transform[currentState][blankInput]
        default:
            currentState = transform[currentState][invalidInput]
        }
        if currentState == invalidState {
            return false
        }
    }
	// Finally, once the string is complete, we make sure we can transform
    // to the done state. This is to prevent strings such as "3.4e" from
	// passing because they are valid up until the very end except for the
	// fact the dangling e makes it invalid. Basically, not all states can
	// transition to "complete" and we check that here.
    currentState = transform[currentState][doneInput]
    if currentState == invalidState {
        return false
    }
    return true
}
