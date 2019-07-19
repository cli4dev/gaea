package vue

const Build = `
#!/usr/bin/env bash
npm run build
go build
mv ./{{.projectName}} ./dist
cd dist
./{{.projectName}} install -r zk://192.168.0.109 -c {{.projectName}}
./{{.projectName}} run -r zk://192.168.0.109 -c {{.projectName}}
#cp -r /root/work/img ./static/static/
#scp -r static root@192.168.106.129:/home/ums/ums
`

const MenuConf = ``
