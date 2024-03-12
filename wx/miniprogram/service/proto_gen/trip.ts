import {rental} from './rental/rental_pb'
import { Coolcar } from './request'
export namespace TripService{
    export function CreateTrip(req:rental.v1.ICreateTripRequest):Promise<rental.v1.ITripEntity>{
        return Coolcar.sendRequestWithAuthRetry({
                method:'POST',
                path:'/v1/rental/trip/create',
                data:req,
                respUnmarshaller:rental.v1.TripEntity.fromObject
            },)
    }
    export function GetTrip(id:string):Promise<rental.v1.ITrip>{
        return Coolcar.sendRequestWithAuthRetry({
            method:'GET',
            path:`v1/rental/trip/get/${encodeURIComponent(id)}`,
            respUnmarshaller:rental.v1.Trip.fromObject
        })
    }
    export function GetTrips(s?:rental.v1.TripStatus):Promise<rental.v1.IGetTripsReponse>{
        let path='v1/rental/trips'
        if (s){
            path+=`?status=${s}`
        }
        return Coolcar.sendRequestWithAuthRetry({
            method:'GET',
            path:path,
            respUnmarshaller:rental.v1.GetTripsReponse.fromObject
        })
    }
}