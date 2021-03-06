PGDMP         .                w         	   myprofile     11.6 (Ubuntu 11.6-1.pgdg18.04+1)     11.6 (Ubuntu 11.6-1.pgdg18.04+1)     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                       false            �           1262    17419 	   myprofile    DATABASE     {   CREATE DATABASE myprofile WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';
    DROP DATABASE myprofile;
          	   myprofile    false            �            1259    17448    bl_comment_h    TABLE     >  CREATE TABLE public.bl_comment_h (
    id integer NOT NULL,
    content text,
    author character varying(45),
    postid integer,
    status integer,
    created_by character varying(45),
    created_at timestamp without time zone,
    updated_by character varying(45),
    updated_at timestamp without time zone
);
     DROP TABLE public.bl_comment_h;
       public      	   myprofile    false            �            1259    17446    bl_comment_h_id_seq    SEQUENCE     �   CREATE SEQUENCE public.bl_comment_h_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.bl_comment_h_id_seq;
       public    	   myprofile    false    202            �           0    0    bl_comment_h_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.bl_comment_h_id_seq OWNED BY public.bl_comment_h.id;
            public    	   myprofile    false    201            �            1259    17437 	   bl_post_h    TABLE     H  CREATE TABLE public.bl_post_h (
    id integer NOT NULL,
    title character varying(45),
    content text,
    author character varying(45),
    status integer,
    created_by character varying(45),
    created_at timestamp without time zone,
    updated_by character varying(45),
    updated_at timestamp without time zone
);
    DROP TABLE public.bl_post_h;
       public      	   myprofile    false            �            1259    17435    bl_post_h_id_seq    SEQUENCE     �   CREATE SEQUENCE public.bl_post_h_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.bl_post_h_id_seq;
       public    	   myprofile    false    200            �           0    0    bl_post_h_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.bl_post_h_id_seq OWNED BY public.bl_post_h.id;
            public    	   myprofile    false    199            �            1259    17431    schema_migration    TABLE     U   CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);
 $   DROP TABLE public.schema_migration;
       public      	   myprofile    false            �            1259    17422 	   xx_user_h    TABLE     �  CREATE TABLE public.xx_user_h (
    id integer NOT NULL,
    name character varying(45),
    email character varying(100),
    typeid integer,
    roleid integer,
    username character varying(45),
    password character varying(255),
    api_token character varying(100),
    remember_token character varying(100),
    status integer,
    created_by character varying(45),
    created_at timestamp without time zone,
    updated_by character varying(45),
    updated_at timestamp without time zone
);
    DROP TABLE public.xx_user_h;
       public      	   myprofile    false            �            1259    17420    xx_user_h_id_seq    SEQUENCE     �   CREATE SEQUENCE public.xx_user_h_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.xx_user_h_id_seq;
       public    	   myprofile    false    197            �           0    0    xx_user_h_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.xx_user_h_id_seq OWNED BY public.xx_user_h.id;
            public    	   myprofile    false    196                       2604    17451    bl_comment_h id    DEFAULT     r   ALTER TABLE ONLY public.bl_comment_h ALTER COLUMN id SET DEFAULT nextval('public.bl_comment_h_id_seq'::regclass);
 >   ALTER TABLE public.bl_comment_h ALTER COLUMN id DROP DEFAULT;
       public    	   myprofile    false    201    202    202                       2604    17440    bl_post_h id    DEFAULT     l   ALTER TABLE ONLY public.bl_post_h ALTER COLUMN id SET DEFAULT nextval('public.bl_post_h_id_seq'::regclass);
 ;   ALTER TABLE public.bl_post_h ALTER COLUMN id DROP DEFAULT;
       public    	   myprofile    false    199    200    200                       2604    17425    xx_user_h id    DEFAULT     l   ALTER TABLE ONLY public.xx_user_h ALTER COLUMN id SET DEFAULT nextval('public.xx_user_h_id_seq'::regclass);
 ;   ALTER TABLE public.xx_user_h ALTER COLUMN id DROP DEFAULT;
       public    	   myprofile    false    196    197    197            �          0    17448    bl_comment_h 
   TABLE DATA               {   COPY public.bl_comment_h (id, content, author, postid, status, created_by, created_at, updated_by, updated_at) FROM stdin;
    public    	   myprofile    false    202   ,        �          0    17437 	   bl_post_h 
   TABLE DATA               w   COPY public.bl_post_h (id, title, content, author, status, created_by, created_at, updated_by, updated_at) FROM stdin;
    public    	   myprofile    false    200   �        �          0    17431    schema_migration 
   TABLE DATA               3   COPY public.schema_migration (version) FROM stdin;
    public    	   myprofile    false    198   $!       �          0    17422 	   xx_user_h 
   TABLE DATA               �   COPY public.xx_user_h (id, name, email, typeid, roleid, username, password, api_token, remember_token, status, created_by, created_at, updated_by, updated_at) FROM stdin;
    public    	   myprofile    false    197   A!       �           0    0    bl_comment_h_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.bl_comment_h_id_seq', 1, true);
            public    	   myprofile    false    201            �           0    0    bl_post_h_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.bl_post_h_id_seq', 2, true);
            public    	   myprofile    false    199            �           0    0    xx_user_h_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.xx_user_h_id_seq', 3, true);
            public    	   myprofile    false    196                       2606    17456    bl_comment_h bl_comment_h_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.bl_comment_h
    ADD CONSTRAINT bl_comment_h_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.bl_comment_h DROP CONSTRAINT bl_comment_h_pkey;
       public      	   myprofile    false    202                       2606    17445    bl_post_h bl_post_h_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.bl_post_h
    ADD CONSTRAINT bl_post_h_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.bl_post_h DROP CONSTRAINT bl_post_h_pkey;
       public      	   myprofile    false    200                       2606    17430    xx_user_h xx_user_h_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.xx_user_h
    ADD CONSTRAINT xx_user_h_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.xx_user_h DROP CONSTRAINT xx_user_h_pkey;
       public      	   myprofile    false    197            	           1259    17434    schema_migration_version_idx    INDEX     c   CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);
 0   DROP INDEX public.schema_migration_version_idx;
       public      	   myprofile    false    198            �   E   x�3��M�+I,��H-*��+.�4����420��54�50T00�21�21�3554���)ghh����� m+      �   �   x�3�tJ�I�J,Rp��I�KW�%
��)�p�t�\1T.5�25��N�SU2,�����`rӨ�a�$?�3#��<3���Ӏ����P���(�����R��P��@��������P������r#��9H��X����̒+F��� ���      �      x������ � �      �   �   x�}�Ak�@�ϻ��C�&���t��hI��M=x�V�4�D�Ĕ�zi(BA�af`��.m^�c�+������d+	�{b��B�VR5xQ�q��j��C�d�3��]��TQi/�e�t:m�������.��������p�|�d��)'��Խ��C��`��h]�(������,�i���	�*$�K'��
y)cm����C@t�+l@!����ӭM)�[M[\     