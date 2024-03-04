import camelcaseKeys from "camelcase-keys"
import { auth } from "./auth/auth_pb"

export namespace Coolcar{
    const serverAddr="http://localhost:8080"
    const authData={
        token:'',
        expireMs:0
    }
    const AUTH_ERR='AUTH_ERR'
    export interface RequestOption<REQ,RES>{
        method:'PUT'|'POST'|'GET'|'DEL',
        path:string,
        data?:REQ,
        respUnmarshaller:(r:object)=>RES
    }
    interface AuthOption{
        RequestWithToken:boolean
        RetryIfFailed:boolean
    }
    export async function login(){
        if(authData.token&&authData.expireMs>=Date.now()){
            return
        }
        const wxResp=await wxLogin()
        const res= await sendRequest<auth.v1.ILoginRequest,auth.v1.ILoginResponse>({
                path:'/v1/auth/login',
                method:'POST',
                data:{
                    code:wxResp.code
                },
                respUnmarshaller:auth.v1.LoginResponse.fromObject
        },{RequestWithToken:false,RetryIfFailed:false})
        authData.token=res.accessToken!
        authData.expireMs=Date.now()+res.expireIn!*1000
    }

    export async function sendRequestWithAuthRetry<REQ,RES>(o:RequestOption<REQ,RES>,a?:AuthOption) :Promise<RES> {
        const authOpt=a||{
            RequestWithToken:true,
            RetryIfFailed:true
        }
        try{
            await login()
            return sendRequest(o,authOpt)
        }catch(err){
            if(err===AUTH_ERR&&authOpt.RetryIfFailed){
                    authData.token=''
                    authData.expireMs=0
                    return sendRequestWithAuthRetry(o,{
                        RequestWithToken:authOpt.RequestWithToken,
                        RetryIfFailed:false
                    })
            }else{
                throw err
            }
        }
    }
    function wxLogin():Promise<WechatMiniprogram.LoginSuccessCallbackResult>{
        return new Promise((resolve,reject)=>{
            wx.login({
                success:resolve,
                fail:reject
            })
        })
    }
    async function sendRequest<REQ,RES>(o:RequestOption<REQ,RES>,a:AuthOption) :Promise<RES> {
        return new Promise((resolve,reject)=>{
            const header:Record<string,any> ={}
            if(a.RequestWithToken===true){
                if(authData.token&&authData.expireMs>=Date.now()){
                    header.authorization='Bearer '+authData.token
                }else{
                    reject(AUTH_ERR)
                    return
                }
            }
            wx.request({
                url:serverAddr+o.path,
                method:'POST',
                data:o.data as WechatMiniprogram.IAnyObject,
                success:res=>{
                    if(res.statusCode===401){
                        reject(AUTH_ERR)
                    }else if(res.statusCode>=400){
                        reject(res)
                    }else{
                        resolve(o.respUnmarshaller(camelcaseKeys(res.data as object,{deep:true})))
                    }
                },
                header,
                fail:reject
            })
        })
    }
}