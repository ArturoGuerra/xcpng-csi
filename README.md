# XCP-ng CSI Driver

XCP-ng CSI Driver is a Go project that uses the Container Storage Interface to integrate native XCP-ng Volumes in Kubernetes or other CSI capable container orchestrators

## Installation

Use the package manager Helm to install the driver in Kubernetes

```bash
helm install xcpng-csi ./chart --namespace=xcpng-csi
```

## Usage

```yaml
# StorageClass that uses the XCP-ng CSI Driver
apiVesrion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: fast
provisioner: csi.xcpng.ar2ro.io
parameters:
   Datastore: "Optional: Storage Repository Name"
   FSType: "Optional: Filesystem Type ie. ext4 Defaults to ext4"
---
# PersistentVolumeClaim that users the fast StorageClass
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
   name: test-claim
spec:
   accessModes:
     - ReadWriteOnce
   resources:
     requests:
       storage: 5Gi
   storageClassName: fast
---
# Pod that consumes the PersistentVolumeClaim
apiVersion: v1
kind: Pod
metadata:
   name: test-pod
spec:
   volumes:
     - name: config
       persistentVolumeClaim:
         claimName: test-claim
   containers:
     - name: test-pod
       image: nginx:latest
       volumeMounts:
         - name: config
           mountPath: /usr/share/nginx/html
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
