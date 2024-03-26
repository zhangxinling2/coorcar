import { rental } from "./rental/rental_pb";
import { Coolcar } from "./request";


export namespace ProfileService{
    export function GetProfile():Promise<rental.v1.IProfile>{
        return Coolcar.sendRequestWithAuthRetry({
            method:'GET',
            path:'/v1/rental/profile',
            respUnmarshaller:rental.v1.Profile.fromObject,
        })
    }
    export function SubmitProfile(req:rental.v1.IIdentity):Promise<rental.v1.IProfile>{
        return Coolcar.sendRequestWithAuthRetry({
            method:'POST',
            path:'/v1/rental/profile',
            data:req,
            respUnmarshaller:rental.v1.Profile.fromObject,
        })
    }
    export function ClearProfile():Promise<rental.v1.IProfile>{
        return Coolcar.sendRequestWithAuthRetry({
            method:'DELETE',
            path:'/v1/rental/profile',
            respUnmarshaller:rental.v1.Profile.fromObject,
        })
    }
}