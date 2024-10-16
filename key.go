package vksol

import (
	"io"
	"text/template"
)

type G1Affine struct {
	X, Y string
}

type E2 struct {
	A0, A1 string
}

type G2Affine struct {
	X, Y E2
}

type VerifyingKey struct {
	// [α]1, [Kvk]1
	G1 struct {
		Alpha       G1Affine
		Beta, Delta G1Affine   // unused, here for compatibility purposes
		K           []G1Affine // The indexes correspond to the public wires
	}
	// [β]2, [δ]2, [γ]2,
	// -[δ]2, -[γ]2: see proof.Verify() for more details
	G2 struct {
		Beta, Delta, Gamma G2Affine
	}
}

func (vk *VerifyingKey) ExportSolidity(w io.Writer) error {
	helpers := template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}

	tmpl, err := template.New("").Funcs(helpers).Parse(solidityTemplate)
	if err != nil {
		return err
	}

	// execute template
	return tmpl.Execute(w, vk)
}
