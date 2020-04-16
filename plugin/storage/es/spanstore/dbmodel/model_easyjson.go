// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dbmodel

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

func easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel(in *jlexer.Lexer, out *Span) {
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
		case "traceID":
			out.TraceID = TraceID(in.String())
		case "spanID":
			out.SpanID = SpanID(in.String())
		case "parentSpanID":
			out.ParentSpanID = SpanID(in.String())
		case "flags":
			out.Flags = uint32(in.Uint32())
		case "operationName":
			out.OperationName = string(in.String())
		case "references":
			if in.IsNull() {
				in.Skip()
				out.References = nil
			} else {
				in.Delim('[')
				if out.References == nil {
					if !in.IsDelim(']') {
						out.References = make([]Reference, 0, 1)
					} else {
						out.References = []Reference{}
					}
				} else {
					out.References = (out.References)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Reference
					easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel1(in, &v1)
					out.References = append(out.References, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "startTime":
			out.StartTime = uint64(in.Uint64())
		case "startTimeMillis":
			out.StartTimeMillis = uint64(in.Uint64())
		case "duration":
			out.Duration = uint64(in.Uint64())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]KeyValue, 0, 1)
					} else {
						out.Tags = []KeyValue{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v2 KeyValue
					easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel2(in, &v2)
					out.Tags = append(out.Tags, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tag":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Tag = make(map[string]interface{})
				} else {
					out.Tag = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 interface{}
					if m, ok := v3.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v3.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v3 = in.Interface()
					}
					(out.Tag)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "logs":
			if in.IsNull() {
				in.Skip()
				out.Logs = nil
			} else {
				in.Delim('[')
				if out.Logs == nil {
					if !in.IsDelim(']') {
						out.Logs = make([]Log, 0, 2)
					} else {
						out.Logs = []Log{}
					}
				} else {
					out.Logs = (out.Logs)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Log
					easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel3(in, &v4)
					out.Logs = append(out.Logs, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "process":
			easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel4(in, &out.Process)
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
func easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel(out *jwriter.Writer, in Span) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"traceID\":"
		out.RawString(prefix[1:])
		out.String(string(in.TraceID))
	}
	{
		const prefix string = ",\"spanID\":"
		out.RawString(prefix)
		out.String(string(in.SpanID))
	}
	if in.ParentSpanID != "" {
		const prefix string = ",\"parentSpanID\":"
		out.RawString(prefix)
		out.String(string(in.ParentSpanID))
	}
	if in.Flags != 0 {
		const prefix string = ",\"flags\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Flags))
	}
	{
		const prefix string = ",\"operationName\":"
		out.RawString(prefix)
		out.String(string(in.OperationName))
	}
	{
		const prefix string = ",\"references\":"
		out.RawString(prefix)
		if in.References == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.References {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel1(out, v6)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"startTime\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.StartTime))
	}
	{
		const prefix string = ",\"startTimeMillis\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.StartTimeMillis))
	}
	{
		const prefix string = ",\"duration\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Duration))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.Tags {
				if v7 > 0 {
					out.RawByte(',')
				}
				easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel2(out, v8)
			}
			out.RawByte(']')
		}
	}
	if len(in.Tag) != 0 {
		const prefix string = ",\"tag\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v9First := true
			for v9Name, v9Value := range in.Tag {
				if v9First {
					v9First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v9Name))
				out.RawByte(':')
				if m, ok := v9Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v9Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v9Value))
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"logs\":"
		out.RawString(prefix)
		if in.Logs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.Logs {
				if v10 > 0 {
					out.RawByte(',')
				}
				easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel3(out, v11)
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"process\":"
		out.RawString(prefix)
		easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel4(out, in.Process)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Span) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Span) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Span) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Span) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel(l, v)
}
func easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel4(in *jlexer.Lexer, out *Process) {
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
		case "serviceName":
			out.ServiceName = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]KeyValue, 0, 1)
					} else {
						out.Tags = []KeyValue{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v12 KeyValue
					easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel2(in, &v12)
					out.Tags = append(out.Tags, v12)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tag":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Tag = make(map[string]interface{})
				} else {
					out.Tag = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v13 interface{}
					if m, ok := v13.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v13.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v13 = in.Interface()
					}
					(out.Tag)[key] = v13
					in.WantComma()
				}
				in.Delim('}')
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
func easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel4(out *jwriter.Writer, in Process) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"serviceName\":"
		out.RawString(prefix[1:])
		out.String(string(in.ServiceName))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Tags {
				if v14 > 0 {
					out.RawByte(',')
				}
				easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel2(out, v15)
			}
			out.RawByte(']')
		}
	}
	if len(in.Tag) != 0 {
		const prefix string = ",\"tag\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v16First := true
			for v16Name, v16Value := range in.Tag {
				if v16First {
					v16First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v16Name))
				out.RawByte(':')
				if m, ok := v16Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v16Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v16Value))
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}
func easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel3(in *jlexer.Lexer, out *Log) {
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
		case "timestamp":
			out.Timestamp = uint64(in.Uint64())
		case "fields":
			if in.IsNull() {
				in.Skip()
				out.Fields = nil
			} else {
				in.Delim('[')
				if out.Fields == nil {
					if !in.IsDelim(']') {
						out.Fields = make([]KeyValue, 0, 1)
					} else {
						out.Fields = []KeyValue{}
					}
				} else {
					out.Fields = (out.Fields)[:0]
				}
				for !in.IsDelim(']') {
					var v17 KeyValue
					easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel2(in, &v17)
					out.Fields = append(out.Fields, v17)
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
func easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel3(out *jwriter.Writer, in Log) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"timestamp\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Timestamp))
	}
	{
		const prefix string = ",\"fields\":"
		out.RawString(prefix)
		if in.Fields == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v18, v19 := range in.Fields {
				if v18 > 0 {
					out.RawByte(',')
				}
				easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel2(out, v19)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel2(in *jlexer.Lexer, out *KeyValue) {
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
		case "key":
			out.Key = string(in.String())
		case "type":
			out.Type = ValueType(in.String())
		case "value":
			if m, ok := out.Value.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Value.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Value = in.Interface()
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
func easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel2(out *jwriter.Writer, in KeyValue) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"key\":"
		out.RawString(prefix[1:])
		out.String(string(in.Key))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		if m, ok := in.Value.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Value.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Value))
		}
	}
	out.RawByte('}')
}
func easyjsonC80ae7adDecodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel1(in *jlexer.Lexer, out *Reference) {
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
		case "refType":
			out.RefType = ReferenceType(in.String())
		case "traceID":
			out.TraceID = TraceID(in.String())
		case "spanID":
			out.SpanID = SpanID(in.String())
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
func easyjsonC80ae7adEncodeGithubComJaegertracingJaegerPluginStorageEsSpanstoreDbmodel1(out *jwriter.Writer, in Reference) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"refType\":"
		out.RawString(prefix[1:])
		out.String(string(in.RefType))
	}
	{
		const prefix string = ",\"traceID\":"
		out.RawString(prefix)
		out.String(string(in.TraceID))
	}
	{
		const prefix string = ",\"spanID\":"
		out.RawString(prefix)
		out.String(string(in.SpanID))
	}
	out.RawByte('}')
}
