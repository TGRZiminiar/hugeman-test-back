
## Sample Response ใน Postman ผมไม่ได้เอาข้อมูลรูปภาพ base64 เข้ามาด้วยนะครับพอไม่งั้นเดียวไฟล์ postman มันจะเยอะเกินไป
## ถ้า Test โดยใช้ local ไม่ใช้ kube ใน postman ให้ Test ด้วย Port 5000
## ถ้า Test โดยใช้ kube ให้ใช้ postman Port 80


## Start Server Need In Local No Kube

### 1. Start Docker Environment
```
docker compose -f docker-compose.db.yml up -d
```

### 2. Select Environment File To Start Golang
go run main.go ./env/dev/.env.dev

#### Command To Migrate
```
go run ./pkg/database/script/migration.go ./env/dev/.env.dev
```

#### Command To Build Dockerfile
```
docker build -f ./build/Dockerfile -t test-hugeman-go:latest .
docker image tag test-hugeman-go:latest tgrziminiar/test-hugeman-go:v.0.1
```

## Command To Handle Kubectl For Testing In Local Only
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.2/deploy/static/provider/cloud/deploy.yaml

kubectl create configmap app-env --from-file=./env/prod/.env

kubectl apply -f ./build/mongo

kubectl apply -f ./build/service.yml

kubectl apply -f ./build/ingress.yml

kubectl apply -f ./build/deployment.yml
```


# Api -> Get All Todo (Maybe Pagination) : No Params Required
# Api -> Get Single Todo (For Update) : Need Todo Id Params

# Api -> Sorting Todo : Need Text Query
# Api -> Searching Todo : Need Text Query

# Api -> Create Todo : Need Req Body
# Api -> Update Todo : Need Req Body

# APi -> Delete Todo : Need Todo Id Params