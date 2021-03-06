package json

import (
	jenkinsv1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

var (
	orig  *jenkinsv1.App
	clone *jenkinsv1.App
)

func TestCreatePatch(t *testing.T) {
	t.Parallel()
	setUp(t)

	clone.Spec.ExposedServices = []string{"foo", "bar"}
	patch, err := CreatePatch(orig, clone)

	assert.NoError(t, err, "patch creation should be successful ")
	assert.Equal(t, `[{"op":"add","path":"/spec/exposedServices","value":["foo","bar"]}]`, string(patch), "the patch should have been empty")
}

func TestCreatePatchNil(t *testing.T) {
	t.Parallel()
	setUp(t)

	_, err := CreatePatch(nil, clone)
	assert.Error(t, err, "nil should not be allowed")
	assert.Equal(t, "'before' cannot be nil when creating a JSON patch", err.Error(), "wrong error message")

	_, err = CreatePatch(orig, nil)
	assert.Error(t, err, "nil should not be allowed")
	assert.Equal(t, "'after' cannot be nil when creating a JSON patch", err.Error(), "wrong error message")
}

func TestCreatePatchNoDiff(t *testing.T) {
	t.Parallel()
	setUp(t)

	patch, err := CreatePatch(orig, clone)

	assert.NoError(t, err, "patch creation should be successful ")
	assert.Equal(t, "[]", string(patch), "the patch should have been empty")
}

func setUp(t *testing.T) {
	orig = &jenkinsv1.App{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-app",
		},
		Spec: jenkinsv1.AppSpec{},
	}

	clone = orig.DeepCopy()
}
