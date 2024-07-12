<script setup lang="ts">
import { UserFilled } from '@element-plus/icons-vue'

function logout(): void {
  const token = useCookie('token')
  token.value = null
  
  navigateTo('/login', {
    external: true,
  })
}

</script>

<template>
  <div
    class="flex justify-between items-center p-6"
  >
    <div /> <!-- Placeholder -->
    <div 
      class="flex gap-4"
    >
      <el-dropdown
        trigger="click"
      >
        <el-avatar
          shape="square"
          :size="40"
          :icon="UserFilled"
        />
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item
                disabled
              >
                Account settings
              </el-dropdown-item>
              <el-dropdown-item
                divided
                @click="logout()"
              >
                <el-text type="danger">
                  Logout
                </el-text>
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
      </el-dropdown>
    </div>
  </div>
</template>
