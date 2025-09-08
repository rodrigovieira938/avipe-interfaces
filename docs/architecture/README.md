```mermaid
---
title: High Level Architecture
---
graph TD
    Framework[Main Dashboard Framework] --> UserAuth[User Authentication]
    UserAuth --> AppPerm[Application Permissions]
    Framework --> AppManager[App Manager]
    AppManager --> AppApi[App Programable Interface]
    Framework --> PublicApi[Public API]
```