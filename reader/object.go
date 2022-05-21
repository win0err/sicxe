package reader

import (
	"io"
	"strings"

	"github.com/win0err/sicxe/vm"
)

type headerRecord struct {
	ProgramName     string
	StartingAddress vm.Word
	ProgramLength   vm.Word
}

type textRecord struct {
	StartingAddress vm.Word
	CodeLength      vm.Byte
	Code            []vm.Byte
}

type modificationRecord struct {
	FieldAddress vm.Word
	FieldLength  vm.Byte
}

type endRecord struct {
	ExecutableInstructionAddress vm.Word
}

type Object struct {
	HeaderRecord        headerRecord
	TextRecords         []textRecord
	ModificationRecords []modificationRecord
	EndRecord           endRecord
}

func ReadAsObject(rd *ObjReader) (*Object, error) {
	obj := &Object{HeaderRecord: headerRecord{}}

	for {
		sectionType, err := rd.rd.ReadByte()
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}

		switch sectionType {
		case 'H': // header record
			name := make([]byte, 6)
			if _, err := rd.Read(name); err != nil {
				return nil, err
			}
			obj.HeaderRecord.ProgramName = strings.Trim(string(name), " ")

			if obj.HeaderRecord.StartingAddress, err = rd.ReadWord(); err != nil {
				return nil, err
			}

			if obj.HeaderRecord.ProgramLength, err = rd.ReadWord(); err != nil {
				return nil, err
			}

		case 'T': // text records
			record := textRecord{}

			if record.StartingAddress, err = rd.ReadWord(); err != nil {
				return nil, err
			}

			if record.CodeLength, err = rd.ReadByte(); err != nil {
				return nil, err
			}

			if record.Code, err = rd.ReadBytes(int(record.CodeLength)); err != nil {
				return nil, err
			}

			obj.TextRecords = append(obj.TextRecords, record)
		case 'M': // modification record
			record := modificationRecord{}

			if record.FieldAddress, err = rd.ReadWord(); err != nil {
				return nil, err
			}

			if record.FieldLength, err = rd.ReadByte(); err != nil {
				return nil, err
			}

			obj.ModificationRecords = append(obj.ModificationRecords, record)
		case 'E': // end record
			if obj.EndRecord.ExecutableInstructionAddress, err = rd.ReadWord(); err != nil {
				return nil, err
			}
		}

		if err = rd.SkipLine(); err != nil {
			return nil, err
		}
	}

	return obj, nil
}
