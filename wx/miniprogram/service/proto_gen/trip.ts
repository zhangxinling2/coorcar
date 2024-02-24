import {rental} from './rental/rental_pb'
import { Coolcar } from './request'
export namespace TripService{
    export function CreateTrip(req:rental.v1.ICreateTripRequest):Promise<rental.v1.ICreateTripResponse>{
        return new Promise((resolve,reject)=>{
            Coolcar.sendRequestWithAuthRetry({
                method:'POST',
                path:'/v1/trip/create',
                data:req,
                respUnmarshaller:rental.v1.CreateTripResponse.fromObject
            },)
        })
    }
}