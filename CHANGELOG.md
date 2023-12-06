# Semantic Versioning Changelog

# [1.2.0](https://github.com/jotaEspig/chronos/compare/v1.1.0...v1.2.0) (2023-12-06)


### Bug Fixes

* **api:** improved authorization checks ([41f9cd1](https://github.com/jotaEspig/chronos/commit/41f9cd14279c8087ef9d4ba49f06e846a84dfaac))
* **frontend:** somethings ([d7f5233](https://github.com/jotaEspig/chronos/commit/d7f52331febd01492e94e90b2827741a5e8e106a))
* **front:** log request error ([150f757](https://github.com/jotaEspig/chronos/commit/150f757155aac480cdd4c4d739fc272fec671b83))
* **front:** send authorization tokens ([7b1b2c2](https://github.com/jotaEspig/chronos/commit/7b1b2c2dd1b5ca04dc3f07fd0ddacc63a049aa61))
* jwt claims struct type field is uint8 now ([d30acdd](https://github.com/jotaEspig/chronos/commit/d30acdd93dc770ce5c896bddcb8080065bcb9bd3))
* **login.css:** removed dumb margin-top ([8971f3d](https://github.com/jotaEspig/chronos/commit/8971f3d31d0d9e972e16ead77f0b92deb5d8a97f))
* **login:** minor errors ([9cd79c9](https://github.com/jotaEspig/chronos/commit/9cd79c912e39cee46e0999391c6d7c3175c8b287))
* **login:** not showing if user exists or not ([34c4d49](https://github.com/jotaEspig/chronos/commit/34c4d4937c2d50217cd229505c6bdf6f6ab79c20))
* **login:** now claims has correct type ([c65bebc](https://github.com/jotaEspig/chronos/commit/c65bebc431a2dde095f0a345fd7bfb1f4798bd42))
* **main.js:** change week buttons ([cebce5b](https://github.com/jotaEspig/chronos/commit/cebce5b1080f4ef9a1d2e8cc65bf9c13eb75e193))
* page title ([479f472](https://github.com/jotaEspig/chronos/commit/479f47209623cdaa842a00157f0d807373faca2a))
* preventing user to try to set its own type ([6be79f4](https://github.com/jotaEspig/chronos/commit/6be79f4d0aa7e61d882a44c8d02f3bbe811c43eb))
* **schedulingapi:** preventing adding multiples scheduling to the same time ([2c09790](https://github.com/jotaEspig/chronos/commit/2c0979086a0d1c63d6f02d0c90d1aeaf8d75d222))
* **time_dao:** query ([ce7d2d8](https://github.com/jotaEspig/chronos/commit/ce7d2d8bf13a63a3e584a5ab42890c129e0f56dc))
* u -> savedUser ([91ff8b6](https://github.com/jotaEspig/chronos/commit/91ff8b6ad15183a34c74235a2d3368a2e71c7a51))
* **user_dao.go:** don't update password ([f7b83b7](https://github.com/jotaEspig/chronos/commit/f7b83b7005f3813a9f74ac0d4c6751a96cbf09b7))
* **user:** type is a uint8 between 0 and 2 ([3016c3a](https://github.com/jotaEspig/chronos/commit/3016c3ad04e510b953b0566d80bb761c74c19029))
* using jwt secret key as env variable ([2a06cd3](https://github.com/jotaEspig/chronos/commit/2a06cd3ab172e88afe16dd5802a98cff130de79d))


### Features

* add redirect buttons to signup and login ([7ba1012](https://github.com/jotaEspig/chronos/commit/7ba10125045c35faa0ed3e0d8038c124cf6678c8))
* **config:** added jwt middleware ([99299df](https://github.com/jotaEspig/chronos/commit/99299df91b596d29ee054f289bffcc37dc23cf85))
* **create_tables.sql:** added password in user table ([99849db](https://github.com/jotaEspig/chronos/commit/99849dbc8789596bdb6e4e8331262f3bdc266560))
* **create_tables.sql:** added type field ([5af6157](https://github.com/jotaEspig/chronos/commit/5af6157fc0862b700f2550efcff914e2a8528a39))
* **employeeapi:** add authentication ([4cc14d0](https://github.com/jotaEspig/chronos/commit/4cc14d084dd92b1f5d250ea3e37149e114867b58))
* **front:** add admin page ([e5f8bbb](https://github.com/jotaEspig/chronos/commit/e5f8bbb235d399077e0b32071b0fbf82289b8dea))
* **front:** add signup page ([5fc47eb](https://github.com/jotaEspig/chronos/commit/5fc47ebc97182124eaa407c3b4a1644bc0cf523d))
* **front:** added login ([d6a52ba](https://github.com/jotaEspig/chronos/commit/d6a52ba565b49bde08a75356c430f467c3c21434))
* **front:** implemented week selection ([998b580](https://github.com/jotaEspig/chronos/commit/998b580ef5f5e8e7f1b71b9fbc74f485d68a1d8c))
* **front:** style changes ([ea3e5a7](https://github.com/jotaEspig/chronos/commit/ea3e5a798b6f4db3d4eb4c4f243e4fd99fc3a0f9))
* **index.html:** add simple navbar with some bugs yet ([ff60099](https://github.com/jotaEspig/chronos/commit/ff60099491883d5c4cc0db0a097c3b8505de9a2a))
* **jwt:** added JWTClaims ([e4ce347](https://github.com/jotaEspig/chronos/commit/e4ce3478c0457a6291e573416d8cb01be03c0a2a))
* **login-frontend:** added stylization ([e100515](https://github.com/jotaEspig/chronos/commit/e100515039f29cedf6af56dce41c004900094482))
* **login-frontend:** added stylization ([ba0ab67](https://github.com/jotaEspig/chronos/commit/ba0ab67e802e4b5a005f710580193ede1c266f53))
* **login:** added login endpoint ([1170741](https://github.com/jotaEspig/chronos/commit/1170741e6fe234aa33cf6431c0db8b658ee0203f))
* **logout.html:** add logout ([53953f2](https://github.com/jotaEspig/chronos/commit/53953f22da6640a4fbd9df675e0d3eaa1e9ff40d))
* **route:** added middlewares field ([a150e70](https://github.com/jotaEspig/chronos/commit/a150e706b4c9063a6060f51543ad550c6d7a3f96))
* **schedulingapi:** add authentication ([a99a062](https://github.com/jotaEspig/chronos/commit/a99a06295c3f9fabb216b84adcdf1c74e23a01a7))
* **signup-frontend:** add stylization ([55f9df0](https://github.com/jotaEspig/chronos/commit/55f9df097c00748e5aadba2c46d636550d7ea48a))
* **timeapi:** added authentication ([6423028](https://github.com/jotaEspig/chronos/commit/64230284ef522d608bb7cb08552b4347fd513faa))
* **user:** added bcrypt operations ([b8ca8b8](https://github.com/jotaEspig/chronos/commit/b8ca8b857f93a47766c21d21b0a54f1d875f92ca))
* **user:** added password field ([c70a36c](https://github.com/jotaEspig/chronos/commit/c70a36ccc1bba86f5934dd84e9dde2e40b1e9e04))
* **user:** added type field ([75c11d2](https://github.com/jotaEspig/chronos/commit/75c11d24e2541651f3233d434242722c9282c7ba))
* **userapi:** added authentication ([033a7c1](https://github.com/jotaEspig/chronos/commit/033a7c15d1c792285182fd0878ca1bda2c9ad9b4))
* **utils.css:** add horizontal align class ([34ad629](https://github.com/jotaEspig/chronos/commit/34ad629bf6d151ce82400da60cbab218898ee9c0))
* **utils.js:** add redirect function ([58715ee](https://github.com/jotaEspig/chronos/commit/58715eeec365908d1161ecffe547782be1c70cd6))

# [1.1.0](https://github.com/jotaEspig/chronos/compare/v1.0.0...v1.1.0) (2023-10-11)


### Bug Fixes

* **front:** back end merge ([dc65328](https://github.com/jotaEspig/chronos/commit/dc65328aff49c66e0c01185595f47333b6d8429b))
* **front:** fixed week day selection request ([78a5673](https://github.com/jotaEspig/chronos/commit/78a56735b31471f14a56757676ef24ab280e843a))


### Features

* **back:** api route to list open schedules ([4fde6ea](https://github.com/jotaEspig/chronos/commit/4fde6eab34a6253a441f0a6bd74cc528ce1c19d5))
* **front:** get shedules from api ([7a38c3b](https://github.com/jotaEspig/chronos/commit/7a38c3b4abe4941eec2ae0ae44310d5528be0702))

# 1.0.0 (2023-10-11)


### Bug Fixes

* **api:** Sanitizing after parsing json ([98135d1](https://github.com/jotaEspig/chronos/commit/98135d14ccf8e36fba24ec8e6852eee906cec679))
* **clean.sql:** order of deletes ([7dab1dc](https://github.com/jotaEspig/chronos/commit/7dab1dc04b94a0fb4b4999fa925416daa2173497))
* error messages ([9142636](https://github.com/jotaEspig/chronos/commit/91426362b5103f1c99f80277f367d3a8db18b722))
* errors messages ([f2ea642](https://github.com/jotaEspig/chronos/commit/f2ea642a5532eb8b3261d114d0f78c0e63d32374))
* **get_times_by_data.sql:** trying new conditions ([bfec9b2](https://github.com/jotaEspig/chronos/commit/bfec9b2dee4b5608954919f79f40e9dc25ca6a04))
* **get_times_by_date.sql:** improved conditions ([f6f082f](https://github.com/jotaEspig/chronos/commit/f6f082fe6f8e0fb2fd82ebda9ee5ae4144b68b5b))
* **readme:** changed CHRONOS_ROOT_DIT to "./chronos-files" ([7d33862](https://github.com/jotaEspig/chronos/commit/7d338620a6c5de973436b1fa778dad8881a673be))
* **tests/utils.go:** path to clean.sql ([54cd691](https://github.com/jotaEspig/chronos/commit/54cd69121a72edc23cbe79e3f29a3b2db7a2c92a))
* **time_dao.go:** getTimesByDateQuery ([7cf2fcc](https://github.com/jotaEspig/chronos/commit/7cf2fcca561a93ba709a1330f2049604c27aadfa))
* **timeapi/time.go:** function name ([5f094d8](https://github.com/jotaEspig/chronos/commit/5f094d89ea2433071710262a7c8ed6e2a20bb562))
* **user.go:** checking if end is after the start ([1bd80aa](https://github.com/jotaEspig/chronos/commit/1bd80aafc2d86b09c04852fe8217e09dffb09c03))
* **userapi/user.go:** error message ([29e8e03](https://github.com/jotaEspig/chronos/commit/29e8e0333bbf176890a3be5dc2a8e6a9cf5ff1bc))
* **userapi/user.go:** error messages ([89cb5f7](https://github.com/jotaEspig/chronos/commit/89cb5f7c306b8a2fa491ca70f5f6197f6b889d30))
* **userapi/user.go:** return status code 200 ([2b47334](https://github.com/jotaEspig/chronos/commit/2b47334fc3a10279f4285203f5a0d5b2e08d8bc8))
* **userapi/user.go:** using function IsValid to check user validity ([9beb931](https://github.com/jotaEspig/chronos/commit/9beb931b66522c198f6f10d8d671c74d903fe5b4))
* using StrictPolicy ([33069a6](https://github.com/jotaEspig/chronos/commit/33069a661d1e76b4fc175f645f8e3eabc2edf12b))


### Features

* added endpoint for getting time by some date ([d152ce6](https://github.com/jotaEspig/chronos/commit/d152ce68ebc1638217844203c16264c587deca06))
* added error when ID is 0 ([4d12970](https://github.com/jotaEspig/chronos/commit/4d12970152d80a2e2c82f6f87dfddb72781027c7))
* added function Sanitize for each model ([04c4727](https://github.com/jotaEspig/chronos/commit/04c472718e9b3166be7b60e88f782cc794348fcc))
* added function to get schedulings by date ([6b1f8e9](https://github.com/jotaEspig/chronos/commit/6b1f8e909fb615455c1e318cd2284211f42c0e89))
* added get schedulings by date ([40900a1](https://github.com/jotaEspig/chronos/commit/40900a193e5159bdd5fa396c846595962a0d427a))
* added IsValid funcion ([58f4374](https://github.com/jotaEspig/chronos/commit/58f4374353c2af360d3c70092745dceb84734025))
* added time controllers ([66231ab](https://github.com/jotaEspig/chronos/commit/66231ab236aad752a5f9ad952c350afe99990939))
* **api/routes.go:** appending AvailablesRoutes from employeeapi ([34868c8](https://github.com/jotaEspig/chronos/commit/34868c84211c4fe1165f2f81fb150318cb613712))
* **api:** using Sanitize function ([f20abb9](https://github.com/jotaEspig/chronos/commit/f20abb9d4c37732429aca8217fd59e09c43be144))
* **common:** added ReadFile function ([763cc33](https://github.com/jotaEspig/chronos/commit/763cc338f3f503be51c2c2d8ebdadcbee7fa565d))
* **config:** added bluemonday policy ([4bfc9c4](https://github.com/jotaEspig/chronos/commit/4bfc9c4c82ad45233b8168e7d05d18941fda6263))
* **config:** added support to SQL operations ([766f12b](https://github.com/jotaEspig/chronos/commit/766f12b0fc78d48c3ab4a73338a60b0b1654cd2e))
* **create_tables.sql:** added index for scheduling.start ([a45030a](https://github.com/jotaEspig/chronos/commit/a45030ae279705529ba6c16e090ba038ad23b961))
* **create_tables.sql:** added indexes for some fields in time ([93c3644](https://github.com/jotaEspig/chronos/commit/93c36449554b6abf73e379661bafe59b52a07e1c))
* **employee_dao.go:** added CRUD operations for Employee ([2f536f7](https://github.com/jotaEspig/chronos/commit/2f536f7afdee9fae066f4c41ee64b4c6598a2175))
* **employee_dao.go:** added update operation for Employee ([5323aa8](https://github.com/jotaEspig/chronos/commit/5323aa838b4a0f0224a81e3dced0b10493a32402))
* **employeeapi:** added createEmployee controller ([22f19c3](https://github.com/jotaEspig/chronos/commit/22f19c334307e196f17151500d9311d01885b5a5))
* **employeeapi:** added deleteEmployee ([2cdd70e](https://github.com/jotaEspig/chronos/commit/2cdd70e281135e2e476093f02965698e33907773))
* **employeeapi:** added getEmployee ([03194da](https://github.com/jotaEspig/chronos/commit/03194da42d667daa10900fa46210683d0c3aae50))
* **files.go:** when not founding a file it shows an errors ([dba07fb](https://github.com/jotaEspig/chronos/commit/dba07fb768b236eee7a95c4f83c7c87a395abf9d))
* **front-end:** base layout from schedule page ([f41be5a](https://github.com/jotaEspig/chronos/commit/f41be5aac63b7ee75727070d907743267c4414d9))
* **front:** implemented schedule rendering ([cbbbd77](https://github.com/jotaEspig/chronos/commit/cbbbd7726ac73b4aa2280133fcdd501abc1b8193))
* **front:** round border to schedule item ([2378c0a](https://github.com/jotaEspig/chronos/commit/2378c0ad3424c0dc8c05a98e7515114927a4863c))
* **front:** schedule items are now properly centered ([6854ffb](https://github.com/jotaEspig/chronos/commit/6854ffb7c57974d41caac8aa90f61dc771a6d837))
* **front:** schedule start visual and improved schedule positioning ([b92c1ce](https://github.com/jotaEspig/chronos/commit/b92c1ce673eee984dc6a4d4ef94e92d5e9ef1ab7))
* **json.go:** added type JsonMap ([7bdad37](https://github.com/jotaEspig/chronos/commit/7bdad373953f8d4cbedb5b1e299b5a1b085ea264))
* **models:** added base structs ([7dcd0bc](https://github.com/jotaEspig/chronos/commit/7dcd0bc65c828203f314d1c1ad391ab9893c2763))
* **models:** added ToMap functions ([2f6897f](https://github.com/jotaEspig/chronos/commit/2f6897f73c527775318f0cf980b01660afc1991c))
* **scheduling_dao.go:** added DAO operations for scheduling object ([776ff50](https://github.com/jotaEspig/chronos/commit/776ff50202f2a7bb37a1baba430712d584b33703))
* **scheduling.go:** added IsValid function ([4deb6fc](https://github.com/jotaEspig/chronos/commit/4deb6fc9a5a564a8c16447c57e6a6faf19c89aba))
* **schedulingapi:** added controllers for scheduling ([36fce39](https://github.com/jotaEspig/chronos/commit/36fce39756f8c0445699d5aeb01dd8d7b45febf6))
* **static_routes.go:** added static route for img/ ([3703c12](https://github.com/jotaEspig/chronos/commit/3703c1273b50828a7d5d41066bf3f3e309bcf330))
* **static-routes:** added css and js static routes ([43c5ff1](https://github.com/jotaEspig/chronos/commit/43c5ff15dceafd85a9e7abc554f161a35b573133))
* **time_dao.go:** added DAO operations for time object ([3abb065](https://github.com/jotaEspig/chronos/commit/3abb0658f1d5efea393968729bbb3937b5b31474))
* **time_dao.go:** added function to get times after some date ([9198e88](https://github.com/jotaEspig/chronos/commit/9198e88212e73988515acc3a33d0b98aedc89eae))
* **time.go:** added IsValid function ([6cb55e2](https://github.com/jotaEspig/chronos/commit/6cb55e221ae136c47bb873e9f06b7726aa348466))
* **user_dao.go:** added CreateUser and FindUserByUsername ([25809a7](https://github.com/jotaEspig/chronos/commit/25809a7f7f8b1722b52bdb9daaec246a2a7930f2))
* **user_dao.go:** added function to delete user by ID ([702933d](https://github.com/jotaEspig/chronos/commit/702933d661fdc6497c5467037c98dce84c2ab22d))
* **user_dao.go:** added UpdateUser ([5a83d28](https://github.com/jotaEspig/chronos/commit/5a83d285db9c9985a7ad23bee899c8d679d1b229))
* **user.go:** added function to check if a user is valid or not ([b932137](https://github.com/jotaEspig/chronos/commit/b9321375bf6c46c1b20bcb3b93b5cbb54c7f3eee))
* **user:** added FindUserById ([2763c7f](https://github.com/jotaEspig/chronos/commit/2763c7f5866b2d6abd69c1003a6ded52540969a8))
* **userapi:** added createUser api endpoint ([0d742fb](https://github.com/jotaEspig/chronos/commit/0d742fb52b586308c4180ac32795d42455f10de8))
* **userapi:** added deleteUser endpoint ([24b94ea](https://github.com/jotaEspig/chronos/commit/24b94eadb9f5799ea86a70f055c5b8d1a997076a))
* **userapi:** added getUser api endpoint ([a0bed30](https://github.com/jotaEspig/chronos/commit/a0bed30e81cbd80330349ef4bfc23ea31737f212))
* **userapi:** added updateUser api endpoint and other changes ([c822e56](https://github.com/jotaEspig/chronos/commit/c822e56c4fa4b027ecd7cd22562b29710b66ea75))
* when initing DB it creates the tables ([96ee115](https://github.com/jotaEspig/chronos/commit/96ee1155768dca6074efc09976c78cc5ac3b7c44))
* **xss_policy.go:** added 2 types of policy ([8504734](https://github.com/jotaEspig/chronos/commit/85047342da308660f14a10cf44de137dbe9487d1))


### Performance Improvements

* **db.db:** adding a limit for max conns ([d2e42dd](https://github.com/jotaEspig/chronos/commit/d2e42ddc6108d60ae714cbba84a26f62a8dba368))
