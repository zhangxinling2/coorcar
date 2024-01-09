"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.routing = void 0;
var routing;
(function (routing) {
    function driving(o) {
        return `/pages/driving/driving?trip_id=${o.trip_id}`;
    }
    routing.driving = driving;
    function lock(o) {
        return `/pages/lock/lock?car_id=${o.car_id}`;
    }
    routing.lock = lock;
    function register(o) {
        const page = '/pages/register/register';
        if (!o) {
            return page;
        }
        return `${page}?redirect=${encodeURIComponent(o.redirectURL)}`;
    }
    routing.register = register;
    function mytrips() {
        return '/pages/mytrips/mytrips';
    }
    routing.mytrips = mytrips;
})(routing = exports.routing || (exports.routing = {}));
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoicm91dGluZy5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbInJvdXRpbmcudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7O0FBQUEsSUFBaUIsT0FBTyxDQTZCdkI7QUE3QkQsV0FBaUIsT0FBTztJQUlwQixTQUFnQixPQUFPLENBQUMsQ0FBYTtRQUNqQyxPQUFPLGtDQUFrQyxDQUFDLENBQUMsT0FBTyxFQUFFLENBQUE7SUFDeEQsQ0FBQztJQUZlLGVBQU8sVUFFdEIsQ0FBQTtJQUlELFNBQWdCLElBQUksQ0FBQyxDQUFVO1FBQzNCLE9BQU8sMkJBQTJCLENBQUMsQ0FBQyxNQUFNLEVBQUUsQ0FBQTtJQUNoRCxDQUFDO0lBRmUsWUFBSSxPQUVuQixDQUFBO0lBT0QsU0FBZ0IsUUFBUSxDQUFDLENBQWlCO1FBQ3RDLE1BQU0sSUFBSSxHQUFDLDBCQUEwQixDQUFBO1FBQ3JDLElBQUcsQ0FBQyxDQUFDLEVBQUM7WUFDRixPQUFPLElBQUksQ0FBQTtTQUNkO1FBQ0QsT0FBTyxHQUFHLElBQUksYUFBYSxrQkFBa0IsQ0FBQyxDQUFDLENBQUMsV0FBVyxDQUFDLEVBQUUsQ0FBQTtJQUNsRSxDQUFDO0lBTmUsZ0JBQVEsV0FNdkIsQ0FBQTtJQUNELFNBQWdCLE9BQU87UUFDbkIsT0FBTyx3QkFBd0IsQ0FBQTtJQUNuQyxDQUFDO0lBRmUsZUFBTyxVQUV0QixDQUFBO0FBQ0wsQ0FBQyxFQTdCZ0IsT0FBTyxHQUFQLGVBQU8sS0FBUCxlQUFPLFFBNkJ2QiIsInNvdXJjZXNDb250ZW50IjpbImV4cG9ydCBuYW1lc3BhY2Ugcm91dGluZ3tcclxuICAgIGV4cG9ydCBpbnRlcmZhY2UgRHJpdmluZ09wdHN7XHJcbiAgICAgICAgdHJpcF9pZDpzdHJpbmdcclxuICAgIH1cclxuICAgIGV4cG9ydCBmdW5jdGlvbiBkcml2aW5nKG86RHJpdmluZ09wdHMpe1xyXG4gICAgICAgIHJldHVybiBgL3BhZ2VzL2RyaXZpbmcvZHJpdmluZz90cmlwX2lkPSR7by50cmlwX2lkfWBcclxuICAgIH1cclxuICAgIGV4cG9ydCBpbnRlcmZhY2UgTG9ja09wdHN7XHJcbiAgICAgICAgY2FyX2lkOnN0cmluZ1xyXG4gICAgfVxyXG4gICAgZXhwb3J0IGZ1bmN0aW9uIGxvY2sobzpMb2NrT3B0cyl7XHJcbiAgICAgICAgcmV0dXJuIGAvcGFnZXMvbG9jay9sb2NrP2Nhcl9pZD0ke28uY2FyX2lkfWBcclxuICAgIH1cclxuICAgIGV4cG9ydCBpbnRlcmZhY2UgUmVnaXN0ZXJPcHRze1xyXG4gICAgICAgIHJlZGlyZWN0PzpzdHJpbmdcclxuICAgIH1cclxuICAgIGV4cG9ydCBpbnRlcmZhY2UgUmVnaXN0ZXJQYXJhbXN7XHJcbiAgICAgICAgcmVkaXJlY3RVUkw6c3RyaW5nXHJcbiAgICB9XHJcbiAgICBleHBvcnQgZnVuY3Rpb24gcmVnaXN0ZXIobz86UmVnaXN0ZXJQYXJhbXMpe1xyXG4gICAgICAgIGNvbnN0IHBhZ2U9Jy9wYWdlcy9yZWdpc3Rlci9yZWdpc3RlcidcclxuICAgICAgICBpZighbyl7XHJcbiAgICAgICAgICAgIHJldHVybiBwYWdlXHJcbiAgICAgICAgfVxyXG4gICAgICAgIHJldHVybiBgJHtwYWdlfT9yZWRpcmVjdD0ke2VuY29kZVVSSUNvbXBvbmVudChvLnJlZGlyZWN0VVJMKX1gXHJcbiAgICB9XHJcbiAgICBleHBvcnQgZnVuY3Rpb24gbXl0cmlwcygpe1xyXG4gICAgICAgIHJldHVybiAnL3BhZ2VzL215dHJpcHMvbXl0cmlwcydcclxuICAgIH1cclxufSJdfQ==