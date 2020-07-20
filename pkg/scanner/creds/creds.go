package creds

import batchv1 "k8s.io/api/batch/v1"

type CredentialAdder interface {
	ShouldAdd() bool
	Add(job *batchv1.Job)
}
