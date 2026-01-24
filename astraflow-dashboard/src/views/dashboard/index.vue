<template>
  <div class="h-full flex flex-col gap-6">
    <!-- Header -->
    <div class="flex justify-between items-center">
        <div>
            <h1 class="text-2xl font-bold text-slate-900 tracking-tight">Executive Dashboard</h1>
            <p class="text-slate-500 mt-1 text-sm">Real-time overview of reimbursement operations.</p>
        </div>
        <div class="text-right">
             <span class="text-xs font-semibold text-slate-400 uppercase tracking-wider">System Status</span>
             <div class="flex items-center gap-2 mt-1">
                 <span class="relative flex h-3 w-3">
                    <span v-if="stats.aiStatus === 'online'" class="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75"></span>
                    <span class="relative inline-flex rounded-full h-3 w-3" :class="stats.aiStatus === 'online' ? 'bg-emerald-500' : 'bg-rose-500'"></span>
                 </span>
                 <span class="text-sm font-medium" :class="stats.aiStatus === 'online' ? 'text-slate-700' : 'text-rose-600'">
                    AI Worker {{ stats.aiStatus === 'online' ? 'Online' : 'Offline' }}
                 </span>
             </div>
        </div>
    </div>

    <!-- Top Row: Metrics -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <!-- Card 1: Pending -->
        <div class="bg-white rounded-xl p-6 shadow-sm border border-slate-200 flex flex-col justify-between h-32">
            <div class="flex justify-between items-start">
                <div class="p-2 bg-amber-50 rounded-lg text-amber-600">
                    <el-icon :size="20"><Timer /></el-icon>
                </div>
                <span class="text-xs font-bold text-amber-600 bg-amber-50 px-2 py-1 rounded-full">Action Req.</span>
            </div>
            <div>
                <div class="text-3xl font-bold text-slate-900">{{ stats.pendingCount }}</div>
                <div class="text-sm text-slate-500 font-medium">Pending Approval</div>
            </div>
        </div>

        <!-- Card 2: Processed Today -->
         <div class="bg-white rounded-xl p-6 shadow-sm border border-slate-200 flex flex-col justify-between h-32">
            <div class="flex justify-between items-start">
                <div class="p-2 bg-blue-50 rounded-lg text-blue-600">
                    <el-icon :size="20"><DocumentChecked /></el-icon>
                </div>
            </div>
            <div>
                <div class="text-3xl font-bold text-slate-900">{{ stats.todayCount }}</div>
                <div class="text-sm text-slate-500 font-medium">Processed Today</div>
            </div>
        </div>

        <!-- Card 3: Total Payout -->
         <div class="bg-white rounded-xl p-6 shadow-sm border border-slate-200 flex flex-col justify-between h-32">
            <div class="flex justify-between items-start">
                <div class="p-2 bg-emerald-50 rounded-lg text-emerald-600">
                    <el-icon :size="20"><Money /></el-icon>
                </div>
            </div>
            <div>
                <div class="text-3xl font-bold text-slate-900">${{ stats.totalAmount.toLocaleString() }}</div>
                <div class="text-sm text-slate-500 font-medium">Total Disbursed</div>
            </div>
        </div>

        <!-- Card 4: AI Accuracy -->
         <div class="bg-white rounded-xl p-6 shadow-sm border border-slate-200 flex flex-col justify-between h-32">
            <div class="flex justify-between items-start">
                <div class="p-2 bg-indigo-50 rounded-lg text-indigo-600">
                    <el-icon :size="20"><Aim /></el-icon>
                </div>
                 <span class="text-xs font-bold text-emerald-600 bg-emerald-50 px-2 py-1 rounded-full">+2%</span>
            </div>
            <div>
                <div class="text-3xl font-bold text-slate-900">{{ stats.aiAccuracy }}%</div>
                <div class="text-sm text-slate-500 font-medium">AI Accuracy Rate</div>
            </div>
        </div>
    </div>

    <!-- Middle Row -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- System Health Panel -->
        <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-6 lg:col-span-1">
            <h3 class="text-lg font-bold text-slate-900 mb-6">System Health</h3>
            
            <div class="space-y-6">
                <!-- RabbitMQ Status -->
                <div>
                    <div class="flex justify-between items-center mb-2">
                        <span class="text-sm font-medium text-slate-600">Message Queue Depth</span>
                        <span class="text-xs font-mono bg-slate-100 px-2 py-0.5 rounded text-slate-600">{{ stats.rabbitMqStatus.queueDepth }} msgs</span>
                    </div>
                    <!-- CSS Bar Chart / Progress -->
                    <div class="w-full bg-slate-100 rounded-full h-2.5 overflow-hidden">
                         <div class="bg-indigo-600 h-2.5 rounded-full transition-all duration-500" :style="{ width: Math.min(stats.rabbitMqStatus.queueDepth * 5, 100) + '%' }"></div>
                    </div>
                    <p class="text-xs text-slate-400 mt-2">Optimal range: 0-20 pending tasks.</p>
                </div>

                <div class="h-px bg-slate-100"></div>

                <!-- Services List -->
                <div class="space-y-3">
                     <div class="flex items-center justify-between">
                        <div class="flex items-center gap-2">
                            <div class="w-2 h-2 rounded-full bg-emerald-500"></div>
                            <span class="text-sm text-slate-700">API Gateway</span>
                        </div>
                        <span class="text-xs text-emerald-600 bg-emerald-50 px-2 py-0.5 rounded font-medium">Healthy</span>
                    </div>
                    <div class="flex items-center justify-between">
                        <div class="flex items-center gap-2">
                            <div class="w-2 h-2 rounded-full bg-emerald-500"></div>
                            <span class="text-sm text-slate-700">Database (MySQL)</span>
                        </div>
                        <span class="text-xs text-emerald-600 bg-emerald-50 px-2 py-0.5 rounded font-medium">Healthy</span>
                    </div>
                     <div class="flex items-center justify-between">
                        <div class="flex items-center gap-2">
                            <div class="w-2 h-2 rounded-full" :class="stats.aiStatus === 'online' ? 'bg-emerald-500' : 'bg-rose-500'"></div>
                            <span class="text-sm text-slate-700">OCR Engine</span>
                        </div>
                        <span class="text-xs px-2 py-0.5 rounded font-medium" :class="stats.aiStatus === 'online' ? 'text-emerald-600 bg-emerald-50' : 'text-rose-600 bg-rose-50'">
                            {{ stats.aiStatus === 'online' ? 'Active' : 'Degraded' }}
                        </span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Recent Activity -->
        <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-0 lg:col-span-2 flex flex-col">
            <div class="p-6 border-b border-slate-100">
                <h3 class="text-lg font-bold text-slate-900">Recent Activity</h3>
            </div>
            <div class="flex-1 overflow-auto p-0">
                <table class="w-full text-left text-sm">
                    <thead class="bg-slate-50 text-slate-500 font-medium">
                        <tr>
                            <th class="px-6 py-3">User</th>
                            <th class="px-6 py-3">Action</th>
                            <th class="px-6 py-3">Time</th>
                            <th class="px-6 py-3">Status</th>
                        </tr>
                    </thead>
                    <tbody class="divide-y divide-slate-100">
                        <tr v-for="(log, i) in activityLogs" :key="i" class="hover:bg-slate-50 transition-colors">
                            <td class="px-6 py-4 font-medium text-slate-900">{{ log.user }}</td>
                            <td class="px-6 py-4 text-slate-600">{{ log.action }}</td>
                            <td class="px-6 py-4 text-slate-400">{{ log.time }}</td>
                            <td class="px-6 py-4">
                                <span class="inline-flex items-center gap-1.5 px-2.5 py-0.5 rounded-full text-xs font-medium" 
                                    :class="log.status === 'success' ? 'bg-emerald-50 text-emerald-700' : 'bg-amber-50 text-amber-700'">
                                    {{ log.status === 'success' ? 'Completed' : 'Pending' }}
                                </span>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
    
    <!-- Bottom Row: Chart Placeholder -->
    <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-6 flex-1 min-h-[300px]">
        <div class="flex justify-between items-center mb-6">
            <h3 class="text-lg font-bold text-slate-900">Monthly Expense Trend</h3>
            <el-select placeholder="This Year" size="small" class="!w-32">
                <el-option label="2023" value="2023" />
            </el-select>
        </div>
        
        <!-- CSS-only Bar Chart Placeholder -->
        <div class="h-64 flex items-end justify-between gap-4 px-4">
            <div v-for="h in [40, 65, 45, 80, 55, 90, 70, 60, 75, 50, 85, 95]" :key="h" class="w-full bg-indigo-50 hover:bg-indigo-100 rounded-t-sm relative group transition-all">
                <div class="absolute bottom-0 w-full bg-indigo-500 rounded-t-sm transition-all duration-700 ease-out group-hover:bg-indigo-600" :style="{ height: h + '%' }"></div>
                 <!-- Tooltip -->
                <div class="opacity-0 group-hover:opacity-100 absolute -top-10 left-1/2 -translate-x-1/2 bg-slate-900 text-white text-xs px-2 py-1 rounded transition-opacity pointer-events-none">
                    ${{ h * 120 }}
                </div>
            </div>
        </div>
        <div class="flex justify-between mt-4 text-xs text-slate-400 uppercase tracking-wider font-semibold px-4">
            <span v-for="m in ['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec']" :key="m">{{ m }}</span>
        </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Timer, DocumentChecked, Money, Aim } from '@element-plus/icons-vue'
import { getStats } from '../../api/dashboard'
import type { DashboardStats } from '../../api/dashboard'

const stats = ref<DashboardStats>({
    pendingCount: 0,
    todayCount: 0,
    totalAmount: 0,
    aiAccuracy: 0,
    aiStatus: 'offline',
    rabbitMqStatus: { connected: false, queueDepth: 0 }
})

const activityLogs = ref([
    { user: 'Admin', action: 'Approved Batch #9921', time: '10 mins ago', status: 'success' },
    { user: 'System (AI)', action: 'Processed Invoice TRC-221', time: '15 mins ago', status: 'success' },
    { user: 'John Doe', action: 'Uploaded Invoice', time: '22 mins ago', status: 'pending' },
    { user: 'Sarah Smith', action: 'Uploaded Invoice', time: '1 hour ago', status: 'pending' },
    { user: 'Admin', action: 'System Backup', time: '3 hours ago', status: 'success' },
])

const loadStats = async () => {
    try {
        // const res = await getStats()
        // stats.value = res.data
        
        // Mock
        await new Promise(r => setTimeout(r, 500))
        stats.value = {
            pendingCount: 12,
            todayCount: 145,
            totalAmount: 42850.00,
            aiAccuracy: 94.2,
            aiStatus: 'online',
            rabbitMqStatus: { connected: true, queueDepth: 4 }
        }
    } catch (e) {
        console.error(e)
    }
}

onMounted(() => {
    loadStats()
})
</script>