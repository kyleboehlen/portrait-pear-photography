import {useRoute, useRouter} from 'vue-router';
import { computed } from 'vue';

export const useAdminRoutes = () => {
    const route = useRoute();
    const router = useRouter();

    const entity = computed(() => route.query.entity ?? "");
    const action = computed(() => route.query.action ?? "");
    const selector = computed(() => route.query.selector ?? "");

    const setEntity = (newEntity) => {
        setAction("").then(
            () => router.replace({
                query: {
                    ...route.query ?? {},
                    entity: newEntity
                }
            })
        );
    };

    const setAction = (newAction) => {
        return setSelector("").then(
            () => router.replace({
                query: {
                    ...route.query ?? {},
                    action: newAction
                }
            })
        );
    };

    const setSelector = (newSelector) => {
        return router.replace({
            query: {
                ...route.query ?? {},
                selector: newSelector
            }
        })
    };

    return { entity, action, selector, setEntity, setAction, setSelector };
};