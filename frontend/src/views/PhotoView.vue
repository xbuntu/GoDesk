<template>
  <div class="photo">
    <div id="viewer"></div>
  </div>
</template>

<script>
// @ is an alias to /src
import {Viewer} from '@photo-sphere-viewer/core';
import '@photo-sphere-viewer/core/index.css'
import '@photo-sphere-viewer/gallery-plugin/index.css'
import sphere from '../assets/photo/sphere.jpg'
import biscayne from '../assets/photo/biscayne.jpg'
import loader from '../assets/photo/loader.gif'
import {GalleryPlugin} from "@photo-sphere-viewer/gallery-plugin";

export default {
  name: 'PhotoView',
  data() {
    return {}
  },
  created() {

  },
  mounted() {
    const viewer = new Viewer({
      container: 'viewer',
      panorama: sphere,
      caption: '第一张',
      loadingImg: loader,
      touchmoveTwoFingers: true,
      mousewheelCtrlKey: true,

      plugins: [
        [GalleryPlugin, {
          visibleOnLoad: true,
        }],
      ],
    });

    const gallery = viewer.getPlugin(GalleryPlugin);

    gallery.setItems([
      {
        id: 'sphere',
        panorama: sphere,
        thumbnail: sphere,
        name: "第一张",
        options: {
          caption: '第一张',
        },
      },
      {
        id: 'key-biscayne',
        panorama: biscayne,
        thumbnail: biscayne,
        name: '第二张',
        options: {
          caption: '第二张',
        },
      },
    ]);
  },
  methods: {}
}
</script>

<style lang="scss" scoped>
.photo {
  padding: 10px;

  #viewer {
    width: calc(100vw - 160px);
    height: calc(100vh - 20px);
  }
}
</style>
