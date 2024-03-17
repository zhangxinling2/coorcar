export const formatTime = (date: Date) => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  return (
    [year, month, day].map(formatNumber).join('/') +
    ' ' +
    [hour, minute, second].map(formatNumber).join(':')
  )
}

const formatNumber = (n: number) => {
  const s = n.toString()
  return s[1] ? s : '0' + s
}
export const avatarUrlKey="avatar_url"
export function getSetting():Promise<WechatMiniprogram.GetSettingSuccessCallbackResult>{
  return new Promise((resolve,reject)=>{
    wx.getSetting({
      //success:res=>resolve(res),
      //res=>resolve(res)这是个函数，resolve本身也是个函数，res=>resolve(res)意思与直接resolve相同，reject同理
      success:resolve,
      fail:reject,
    })
  })
}
export function getUserInfo():Promise<WechatMiniprogram.GetUserInfoSuccessCallbackResult>{
  return new Promise((resolve, reject) => {
    wx.getUserInfo({
      success:resolve,
      fail:reject,
    })
  })
}