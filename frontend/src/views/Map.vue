<template>
    <div id="container"></div>
</template>

<script setup>
import { onMounted, onUnmounted } from "vue";
import AMapLoader from "@amap/amap-jsapi-loader";

let map = null;

onMounted(() => {
    window._AMapSecurityConfig = {
        securityJsCode: "90c1989e9d5c9f755b43f83038182ecc",
    };
    AMapLoader.load({
        key: "b847ae423b0ae1168f433ff791f08815",
        version: "2.0",
        plugins: ["AMap.Scale", 'AMap.MoveAnimation', 'AMap.Driving', 'AMap.AutoComplete', 'AMap.Geolocation'],
    })
        .then((AMap) => {
            map = new AMap.Map("container", {
                viewMode: "3D",
                zoom: 11,
                center: [116.397428, 39.90923],
            });
            const getCurrentLocation = () => {
                return new Promise((resolve, reject) => {
                    const geolocation = new AMap.Geolocation({
                        enableHighAccuracy: true, // 高精度定位
                        timeout: 10000,           // 超时时间
                        showButton: false,        // 不显示高德自带的定位按钮
                    });

                    geolocation.getCurrentPosition((status, result) => {
                        if (status === 'complete') {
                            // result.position 包含 lng 和 lat
                            resolve({
                                lng: result.position.lng,
                                lat: result.position.lat,
                            });
                        } else {
                            reject(result);
                        }
                    });
                });
            };
            console.log(getCurrentLocation())
        })
        .catch((e) => {
            console.log(e);
        });
});

onUnmounted(() => {
    map?.destroy();
});
</script>

<style scoped>
#container {
    padding: 0px;
    margin: 0px;
    width: 100%;
    height: 800px;
}
</style>
