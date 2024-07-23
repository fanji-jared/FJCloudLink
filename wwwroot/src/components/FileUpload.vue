<template>
  <n-upload
    multiple
    directory-dnd
    :max="5"
    v-bind:directory="true"
    @change="handleChange"
  >
    <n-upload-dragger>
      <div style="margin-bottom: 12px">
        <n-icon size="48" :depth="3">
          <ArchiveIcon />
        </n-icon>
      </div>
      <n-text style="font-size: 16px"> 点击或者拖动文件到该区域来上传 </n-text>
      <n-p depth="3" style="margin: 8px 0 0 0">
        请不要上传敏感数据，比如你的银行卡号和密码，信用卡号有效期和安全码
      </n-p>
    </n-upload-dragger>
  </n-upload>
</template>

<script>
import { defineComponent } from "vue";
import { ArchiveOutline as ArchiveIcon } from "@vicons/ionicons5";
//
// import { NUpload, NUploadDragger, NIcon, NText} from 'naive-ui';

export default defineComponent({
  components: {
    ArchiveIcon,
    // NUpload,
    // NUploadDragger,
    // NIcon,
    // NText,
  },
});
</script>