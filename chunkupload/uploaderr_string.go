// Code generated by "stringer -type uploadErr -linecomment"; DO NOT EDIT.

package chunkupload

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrDirInvalid-1]
	_ = x[ErrEvalAbsPath-2]
}

const _uploadErr_name = "uploader: dir invaliduploader: cannot evaluate absolute path"

var _uploadErr_index = [...]uint8{0, 21, 60}

func (i uploadErr) String() string {
	i -= 1
	if i < 0 || i >= uploadErr(len(_uploadErr_index)-1) {
		return "uploadErr(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _uploadErr_name[_uploadErr_index[i]:_uploadErr_index[i+1]]
}
