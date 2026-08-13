// Package cloudkitty provides CloudKitty service configuration and management utilities
package cloudkitty

import (
	corev1 "k8s.io/api/core/v1"
)

var (
	config0440AccessMode int32 = 0440
	certMode             int32 = 0400
)

// GetVolumes - service volumes
func GetVolumes(name string) []corev1.Volume {
	return []corev1.Volume{
		{
			Name: "config-data",
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					DefaultMode: &config0440AccessMode,
					SecretName:  name + "-config-data",
				},
			},
		}, {
			Name: "certs",
			VolumeSource: corev1.VolumeSource{
				Projected: &corev1.ProjectedVolumeSource{
					Sources: []corev1.VolumeProjection{
						{
							Secret: &corev1.SecretProjection{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: ClientCertSecretName,
								},
							},
						}, {
							ConfigMap: &corev1.ConfigMapProjection{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: name + "-lokistack-gateway-ca-bundle",
								},
							},
						},
					},
					DefaultMode: &certMode,
				},
			},
		},
	}
}

// GetVolumeMounts - VolumeMounts shared by api and proc (metrics + loki certs).
// The base cloudkitty.conf is NOT included here — api/proc get it via
// config-data-custom at conf.d/00-cloudkitty.conf; dbsync/storageinit mount
// it directly in their own inline volumeMount lists.
func GetVolumeMounts() []corev1.VolumeMount {
	return []corev1.VolumeMount{
		{
			Name:      "config-data",
			MountPath: "/etc/cloudkitty/metrics.yaml",
			SubPath:   "metrics.yaml",
			ReadOnly:  true,
		},
		{
			Name:      "certs",
			MountPath: "/etc/cloudkitty/certs",
			ReadOnly:  true,
		},
	}
}
