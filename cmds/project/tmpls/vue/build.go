package vue

const Build = `
#!/usr/bin/env bash
npm run build
#go install
#cp ${GOPATH}/bin/ums ./dist
#cd dist
#cp -r /root/work/img ./static/static/
#scp -r static root@192.168.106.129:/home/ums/ums
`
