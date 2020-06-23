// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson5a72dc82DecodeMainModels(in *jlexer.Lexer, out *VoteDB) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = int(in.Int())
		case "Nickname":
			out.Nickname = string(in.String())
		case "Thread_id":
			out.Thread_id = int(in.Int())
		case "Voice":
			out.Voice = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5a72dc82EncodeMainModels(out *jwriter.Writer, in VoteDB) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"Nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"Thread_id\":"
		out.RawString(prefix)
		out.Int(int(in.Thread_id))
	}
	{
		const prefix string = ",\"Voice\":"
		out.RawString(prefix)
		out.Int(int(in.Voice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VoteDB) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeMainModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VoteDB) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeMainModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VoteDB) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeMainModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VoteDB) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeMainModels(l, v)
}
func easyjson5a72dc82DecodeMainModels1(in *jlexer.Lexer, out *Vote) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "nickname":
			out.Nickname = string(in.String())
		case "voice":
			out.Voice = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5a72dc82EncodeMainModels1(out *jwriter.Writer, in Vote) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix[1:])
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"voice\":"
		out.RawString(prefix)
		out.Int(int(in.Voice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Vote) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeMainModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Vote) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeMainModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Vote) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeMainModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Vote) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeMainModels1(l, v)
}
func easyjson5a72dc82DecodeMainModels2(in *jlexer.Lexer, out *PostUpdate) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message":
			if in.IsNull() {
				in.Skip()
				out.Message = nil
			} else {
				if out.Message == nil {
					out.Message = new(string)
				}
				*out.Message = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5a72dc82EncodeMainModels2(out *jwriter.Writer, in PostUpdate) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix[1:])
		if in.Message == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Message))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostUpdate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeMainModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUpdate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeMainModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUpdate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeMainModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUpdate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeMainModels2(l, v)
}
func easyjson5a72dc82DecodeMainModels3(in *jlexer.Lexer, out *PostDetails) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "author":
			if in.IsNull() {
				in.Skip()
				out.AuthorDetails = nil
			} else {
				if out.AuthorDetails == nil {
					out.AuthorDetails = new(User)
				}
				(*out.AuthorDetails).UnmarshalEasyJSON(in)
			}
		case "forum":
			if in.IsNull() {
				in.Skip()
				out.ForumDetails = nil
			} else {
				if out.ForumDetails == nil {
					out.ForumDetails = new(Forum)
				}
				(*out.ForumDetails).UnmarshalEasyJSON(in)
			}
		case "post":
			if in.IsNull() {
				in.Skip()
				out.PostDetails = nil
			} else {
				if out.PostDetails == nil {
					out.PostDetails = new(Post)
				}
				(*out.PostDetails).UnmarshalEasyJSON(in)
			}
		case "thread":
			if in.IsNull() {
				in.Skip()
				out.ThreadDetails = nil
			} else {
				if out.ThreadDetails == nil {
					out.ThreadDetails = new(Thread)
				}
				(*out.ThreadDetails).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5a72dc82EncodeMainModels3(out *jwriter.Writer, in PostDetails) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AuthorDetails != nil {
		const prefix string = ",\"author\":"
		first = false
		out.RawString(prefix[1:])
		(*in.AuthorDetails).MarshalEasyJSON(out)
	}
	if in.ForumDetails != nil {
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.ForumDetails).MarshalEasyJSON(out)
	}
	if in.PostDetails != nil {
		const prefix string = ",\"post\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.PostDetails).MarshalEasyJSON(out)
	}
	if in.ThreadDetails != nil {
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.ThreadDetails).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostDetails) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeMainModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostDetails) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeMainModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostDetails) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeMainModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostDetails) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeMainModels3(l, v)
}
func easyjson5a72dc82DecodeMainModels4(in *jlexer.Lexer, out *PostArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostArr, 0, 8)
			} else {
				*out = PostArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *Post
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(Post)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5a72dc82EncodeMainModels4(out *jwriter.Writer, in PostArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v PostArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeMainModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeMainModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeMainModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeMainModels4(l, v)
}
func easyjson5a72dc82DecodeMainModels5(in *jlexer.Lexer, out *Post) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "author":
			out.User_nick = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "forum":
			out.Forum_slug = string(in.String())
		case "thread":
			out.Thread_id = int(in.Int())
		case "isEdited":
			out.Is_edited = bool(in.Bool())
		case "parent":
			out.Parent = int32(in.Int32())
		case "Parents":
			if in.IsNull() {
				in.Skip()
				out.Parents = nil
			} else {
				in.Delim('[')
				if out.Parents == nil {
					if !in.IsDelim(']') {
						out.Parents = make([]int32, 0, 16)
					} else {
						out.Parents = []int32{}
					}
				} else {
					out.Parents = (out.Parents)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Parents = append(out.Parents, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5a72dc82EncodeMainModels5(out *jwriter.Writer, in Post) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.User_nick))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.Raw((in.Created).MarshalJSON())
	}
	{
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum_slug))
	}
	{
		const prefix string = ",\"thread\":"
		out.RawString(prefix)
		out.Int(int(in.Thread_id))
	}
	{
		const prefix string = ",\"isEdited\":"
		out.RawString(prefix)
		out.Bool(bool(in.Is_edited))
	}
	if in.Parent != 0 {
		const prefix string = ",\"parent\":"
		out.RawString(prefix)
		out.Int32(int32(in.Parent))
	}
	{
		const prefix string = ",\"Parents\":"
		out.RawString(prefix)
		if in.Parents == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Parents {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Post) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeMainModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Post) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeMainModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Post) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeMainModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Post) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeMainModels5(l, v)
}
